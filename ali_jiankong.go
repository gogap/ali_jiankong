package ali_jiankong

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gogap/errors"
	"github.com/parnurzeal/gorequest"
)

const (
	REPORT_TIMEOUT time.Duration = 3 * time.Second

	ALI_JIANKONG_NAMESPACE = "acs/custom/"
	ALI_JIANKONG_URL       = "http://open.cms.aliyun.com/metrics/put"
)

type Dimensions map[string]string

type ReportItem struct {
	MetricName      string     `json:"metricName"`
	MetricValue     string     `json:"value"`
	Dimensions      Dimensions `json:"dimensions"`
	DimensionsOrder []string   `json:"-"`
	Unit            string     `json:"unit"`
	Timestamp       string     `json:"timestamp"`
}

func (p *ReportItem) Serialize() string {
	arrayDimensions := []string{}
	for _, dim := range p.DimensionsOrder {
		if v, exist := p.Dimensions[dim]; exist {
			arrayDimensions = append(arrayDimensions, fmt.Sprintf("\"%s\":\"%s\"", dim, v))
		} else {
			arrayDimensions = append(arrayDimensions, fmt.Sprintf("\"%s\":\"%s\"", dim, ""))
		}
	}

	return fmt.Sprintf("{\"metricName\":\"%s\",\"value\":\"%s\",\"dimensions\":{%s},\"unit\":\"%s\",\"timestamp\":\"%s\"}",
		p.MetricName,
		p.MetricValue,
		strings.Join(arrayDimensions, ","),
		p.Unit,
		p.Timestamp)
}

type AliJianKong struct {
	uid     string
	timeout time.Duration
}

func NewAliJianKong(uid string, timeout time.Duration) *AliJianKong {
	if timeout == 0 {
		timeout = REPORT_TIMEOUT
	}

	return &AliJianKong{
		uid:     uid,
		timeout: timeout,
	}
}

func (p *AliJianKong) SetTimeout(timeout time.Duration) *AliJianKong {
	p.timeout = timeout
	return p
}

func (p *AliJianKong) Report(items ...ReportItem) (err error) {
	timestamp := time.Now().Unix()
	strTimestamp := strconv.Itoa(int(timestamp))

	metrics := []string{}
	for _, item := range items {
		item.Unit = "None"
		item.Timestamp = strTimestamp
		metrics = append(metrics, item.Serialize())
	}

	strMetrics := "[" + strings.Join(metrics, ",") + "]"

	v := url.Values{}
	v.Add("userId", p.uid)
	v.Add("namespace", ALI_JIANKONG_NAMESPACE+p.uid)
	v.Add("metrics", strMetrics)

	resp, body, errs := gorequest.New().Timeout(p.timeout).Post(ALI_JIANKONG_URL).Send(v.Encode()).End()

	if e := errs_to_error(errs); e != nil {
		err = ERR_REQUEST_JIANKONG_SERVER_FAILED.New(errors.Params{"err": e})
		return
	}

	if resp.StatusCode != 200 {
		err = ERR_SEND_JIANKONG_REPORT_FAILED.New(errors.Params{"code": resp.StatusCode, "content": body})
		return
	}

	return
}

func errs_to_error(errs []error) error {
	if errs == nil || len(errs) == 0 {
		return nil
	}

	strErr := ""
	for _, e := range errs {
		strErr += e.Error() + "; "
	}
	return errors.New(strErr)
}

package ali_jiankong

import (
	"github.com/gogap/errors"
)

const (
	ALI_JIANKONG_ERROR_NS = "ALI_JIANKONG"
)

var (
	ERR_SEND_JIANKONG_REPORT_FAILED    = errors.TN(ALI_JIANKONG_ERROR_NS, 1, "code: {{.code}}, content: {{.content}}")
	ERR_REQUEST_JIANKONG_SERVER_FAILED = errors.TN(ALI_JIANKONG_ERROR_NS, 2, "error: {{.err}}")
	ERR_MARSHAL_METRICS_FAILED         = errors.TN(ALI_JIANKONG_ERROR_NS, 3, "marshal mertrics failed, error: {{.err}}")
)

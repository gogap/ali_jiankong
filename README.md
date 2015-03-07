# ali_jiankong

### Usage

```go
package main

import (
	"fmt"

	"github.com/gogap/ali_jiankong"
)

func main() {

	cli := ali_jiankong.NewAliJianKong("191000000000000", ali_jiankong.REPORT_TIMEOUT)

	reportItem := ali_jiankong.ReportItem{
		MetricName:  "test2",
		MetricValue: "1",
		Dimensions:  ali_jiankong.Dimensions{"aaaa": "1", "bbbb": "2", "cccc": "3", "dddd": "4", "eeee": "5"},
	}

	e := cli.Report(reportItem)
	fmt.Println(e)

}
```
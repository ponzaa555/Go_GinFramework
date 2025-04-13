package log

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type logFormatLocal struct {
	TimeStamp    time.Time
	StatusCode   int
	Latency      time.Duration
	ClientIP     string
	Method       string
	Path         string
	ErrorMessage string
	RequestProto string
}

func FormatLogsJson(param gin.LogFormatterParams) string {
	params := &logFormatLocal{
		ClientIP:     param.ClientIP,
		TimeStamp:    param.TimeStamp,
		Method:       param.Method,
		Path:         param.Path,
		RequestProto: param.Request.Proto,
		StatusCode:   param.StatusCode,
		Latency:      param.Latency,
		ErrorMessage: param.ErrorMessage,
	}
	j, _ := json.Marshal(params)
	return string(j)

}

func FormatsLogs(param gin.LogFormatterParams) string {
	return fmt.Sprintf("{ %s - [%s] \"%s %s %s %d %s \"%s\" %s\"} \n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

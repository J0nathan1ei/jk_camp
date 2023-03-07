package logger

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var Log = logrus.New()

func init() {
	// 设置日志格式为 JSON 格式
	Log.Formatter = &logrus.JSONFormatter{}
	// 设置日志级别为 info 以上
	Log.Level = logrus.InfoLevel
	// 设置日志输出为标准输出
	Log.Out = os.Stdout
}

func RecordHttpClient(r *http.Request, uri string) {
	ipAddr := r.RemoteAddr
	Log.WithFields(logrus.Fields{
		"remote_ip":   ipAddr,
		"status_code": http.StatusOK,
	}).Info(uri)
}

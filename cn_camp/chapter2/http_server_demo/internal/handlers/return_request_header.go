package handlers

import (
	"fmt"
	"http_server_demo/internal/logger"
	"net/http"
	"strings"
)

func RequestHeaderHandler(w http.ResponseWriter, r *http.Request) {
	var headerStr string
	// 将 HTTP 请求头转换为字符串
	for k, v := range r.Header {
		vStr := strings.Join(v, " ")
		headerStr += k + ": " + vStr + "\n"
	}
	w.Header().Set("request_header", headerStr)
	fmt.Fprintf(w, "go check your response header!")
	logger.RecordHttpClient(r, "request_header")
}

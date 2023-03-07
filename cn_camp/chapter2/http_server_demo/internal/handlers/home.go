package handlers

import (
	"fmt"
	"http_server_demo/internal/logger"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a http server's home page!")
	logger.RecordHttpClient(r, "home")
}

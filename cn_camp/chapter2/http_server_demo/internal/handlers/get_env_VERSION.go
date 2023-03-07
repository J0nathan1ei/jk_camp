package handlers

import (
	"fmt"
	"http_server_demo/internal/logger"
	"net/http"
	"os"
)

func EnvVersionHandler(w http.ResponseWriter, r *http.Request) {
	if version := os.Getenv("VERSION"); len(version) == 0 {
		fmt.Fprintf(w, "your env VERSION is %s", version)
	} else {
		fmt.Fprintf(w, "your env VERSION is empty")
	}
	logger.RecordHttpClient(r, "env_version")
}

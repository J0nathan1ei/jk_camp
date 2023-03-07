package main

import (
	"http_server_demo/internal/handlers"
	"http_server_demo/internal/logger"
	"net/http"
)

func main() {
	// 绑定路由
	http.HandleFunc("/", handlers.HomeHandler)

	// 接收客户端 request，并将 request 中带的 header 写入 response header
	http.HandleFunc("/get_req_header", handlers.RequestHeaderHandler)

	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	http.HandleFunc("/get_env_version", handlers.EnvVersionHandler)

	// 当访问 localhost/healthz 时，应返回 200
	http.HandleFunc("/healthz", handlers.HealthZHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Log.Errorf(err.Error())
	}
}

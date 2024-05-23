package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// 获取服务器端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口
	}

	// 启动服务器
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

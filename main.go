package main

import (
	"log"
	"net/http"
	"proxy/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 请求结果匹配
	// proxy/query     匹配
	// proxy/invoke    匹配
	// proxy/v1/find   匹配
	// proxy/          匹配
	r.Any("/proxy/*name", controller.ProxyHandler1) //http://127.0.0.1:8080/proxy/test
	r.Any("/api/*name", controller.ProxyHandler2)   //http://127.0.0.1:8080/api/test

	srv := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

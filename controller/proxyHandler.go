package controller

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyHandler1(c *gin.Context) {
	var target = "http://127.0.0.1:8081"
	proxyUrl, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)

	proxy.ServeHTTP(c.Writer, c.Request)
}

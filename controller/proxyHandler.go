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

	// proxy.ModifyResponse = func(resp *http.Response) error {
	// 	// NOTE 修改响应给客户端的response, 或者用于统计响应时间
	// 	return nil
	// }

	// proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
	// 	// 當執行proxy.ModifyResponse拋錯時，就會執行proxy.ErrorHandler
	// }

	proxy.ServeHTTP(c.Writer, c.Request)
}

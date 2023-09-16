package controller

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// sample
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

// 進階處理request
func ProxyHandler2(c *gin.Context) {
	var target = "http://127.0.0.1:8081"
	proxyUrl, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)

	token := "abcd1234"
	apiHeader := "x-api-token"
	originalDirector := proxy.Director         // 先将原本的处理函数缓存
	proxy.Director = func(req *http.Request) { // 重新赋值新的处理函数
		originalDirector(req) // 执行原本的处理函数

		if req.URL.Query().Get("debug") == "1" { // 故意測試token校驗失敗的response
			token = "test-error"
		}
		req.Header.Set(apiHeader, token) // 增加我们想要对 request 做的操作代码
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

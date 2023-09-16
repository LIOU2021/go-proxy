package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("*path", func(c *gin.Context) {
		path := c.Param("path")
		resp := gin.H{
			"path":   path,
			"host":   c.Request.Host,
			"header": c.Request.Header,
			"url":    c.Request.URL,
		}
		if token := c.GetHeader("x-api-token"); token != "" {
			if token != "abcd1234" {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}

		c.JSON(200, resp)
	})

	srv := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

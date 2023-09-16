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
		c.JSON(200, gin.H{
			"path":   path,
			"host":   c.Request.Host,
			"header": c.Request.Header,
			"url":    c.Request.URL,
		})
	})

	srv := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

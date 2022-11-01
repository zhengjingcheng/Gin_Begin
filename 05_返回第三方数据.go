package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/GetOtherData", func(c *gin.Context) {
		url := "http://www.baidu.com"
		get, err := http.Get(url)
		if err != nil || get.StatusCode != http.StatusOK {
			panic(err)
			return
		}
		body := get.Body
		contentLength := get.ContentLength
		contentType := get.Header.Get("Content-Type")
		//数据写入响应体
		c.DataFromReader(http.StatusOK, contentLength, contentType, body, nil)
	})
	r.Run(":9090")
}

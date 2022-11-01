package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 一般重定向：重定向到外部网络
	r.GET("/redirect1", func(c *gin.Context) {
		//重定向到百度，获取百度对应的数据
		url := "http://www.baidu.com"
		c.Redirect(http.StatusMovedPermanently, url)
	})
	//路由重定向：重定向到具体的路由
	r.GET("/redirect2", func(c *gin.Context) {
		c.Request.URL.Path = "/TestRedirect" //重定向的路由
		r.HandleContext(c)
	})
	//路由
	r.GET("/TestRedirect", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "响应成功",
		})
	})
	r.Run(":9090")
}

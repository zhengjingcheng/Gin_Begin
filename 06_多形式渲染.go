package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//json格式输出
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>hello,GIN框架</b>",
		})
	})
	//原样输出
	r.GET("/someHTML", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>hello,GIN框架</b>",
		})
	})
	//输出xml格式
	r.GET("/somexml", func(c *gin.Context) {
		type Message struct {
			Name string
			Msg  string
			Age  int
		}
		info := Message{}
		info.Name = "阿晓"
		info.Msg = "hello"
		info.Age = 23
		c.XML(http.StatusOK, info)
	})
	//返回yaml形式
	r.GET("/SOMEYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"message": "GIN框架的多形式渲染",
			"status":  200,
		})
	})
	r.Run(":9090")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get", getMsg)
	//r.Run("127.0.0.1:9090")
	r.Run(":9090")
}

func getMsg(c *gin.Context) {
	name := c.Query("name")                 //根据url获取name参数
	c.String(http.StatusOK, "欢迎您:%s", name) //通过string渲染，响应码和显示内容
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "返回信息",
	})
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//传递数据
	r.POST("/post", postMsg)
	r.Run(":9090")
}
func postMsg(c *gin.Context) {
	//name := c.Query("name") //获取Url中的数据
	//方法二读取参数（form是参数值，b代表返回的参数值是否存在）
	form, b := c.GetPostForm("name")
	fmt.Println(form, b)
	//方法一读取参数
	name := c.DefaultPostForm("name", "gin")
	c.JSON(http.StatusOK, "欢迎您："+name)
}

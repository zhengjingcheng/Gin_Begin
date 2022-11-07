package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(AuthMiddleware())
	r.GET("/login", func(c *gin.Context) {
		//获取用户，它是由basucauth中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(http.StatusOK, "登录成功，欢迎您："+user)
	})
	r.Run(":9090")
}
func AuthMiddleware() gin.HandlerFunc {
	//初始化用户
	accounts := gin.Accounts{
		"admin":  "adminpw",
		"system": "systempw",
	}
	//动态添加用户
	accounts["Go"] = "123456789"
	accounts["Gin"] = "123"
	//将用户添加到登录中间件中去
	auth := gin.BasicAuth(accounts)
	return auth
}

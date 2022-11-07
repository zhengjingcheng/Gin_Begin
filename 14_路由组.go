package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResGroup struct {
	Data string
	Path string
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		r := v1.Group("/user") //二级路径
		r.GET("/login", login) //响应请求:/v1/user/login
		//路由分组三级路径
		r2 := r.Group("/showInfo")
		r2.GET("/abstract", abstract) //响应请求：v1/user/showInfo/abstract
		r2.GET("/detail", detail)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/other", other)
	}
	router.Run(":9090")

}

func other(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{
		Data: "other",
		Path: c.Request.URL.Path,
	})
}

func detail(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{
		Data: "detail",
		Path: c.Request.URL.Path,
	})
}

func abstract(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{
		Data: "abstract",
		Path: c.Request.URL.Path,
	})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{
		Data: "login",
		Path: c.Request.URL.Path,
	})
}
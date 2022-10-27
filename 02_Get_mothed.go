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
	name := c.Query("name")
	c.String(http.StatusOK, "欢迎您hhh:%s", name)

}

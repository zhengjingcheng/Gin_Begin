package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/file", fileServer)
	r.Run(":9090")
}

func fileServer(c *gin.Context) {
	path := "E:\\求职\\go_gin"
	fileName := path + c.Query("name")
	fmt.Println(fileName)
	c.File(fileName)
}

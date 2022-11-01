package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("UPLOAD", func(c *gin.Context) {
		file, err := c.FormFile("fileName")
		if err != nil {
			panic(err)
		}
		dst := "E:/求职/go_gin/"
		c.SaveUploadedFile(file, dst+file.Filename)
		c.String(200, fmt.Sprintf("%s上传完成", file.Filename))
	})
	r.Run(":9090")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		//多文件的上传
		form, err := c.MultipartForm() //获取form
		if err != nil {
			c.String(http.StatusBadRequest, "上传文件错误")
		}
		files := form.File["file_key"] //上传的所有文件
		dst := "E:/求职/go_gin/"
		//遍历文件
		for _, file := range files {
			c.SaveUploadedFile(file, dst+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d 个文件上传完成！", len(files)))
	})
	r.Run(":9090")
}

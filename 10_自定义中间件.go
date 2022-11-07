package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default() //默认路由引擎：包括一些常用的中间件
	//r := gin.New() // 没有任何中间件的路由引擎
	r.Use(Middleware())
	r.GET("/middle", func(c *gin.Context) {
		fmt.Println("服务端开始执行")
		fmt.Println("服务端执行开始执行....")
		name := c.Query("name")
		ageStr := c.Query("age")
		age, _ := strconv.Atoi(ageStr)
		log.Println(name, age)
		res := struct {
			Name string `json:"name"` //为结构体添加标签,在json中显示为：name
			Age  int    `json:"age"`
		}{name, age}
		c.JSON(http.StatusOK, res)
	})
	r.Run(":9090")
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件开始执行===")
		name := c.Query("name")
		ageStr := c.Query("age")
		age, err := strconv.Atoi(ageStr) //string转int
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "输入的数据错误，年龄不是整数")
			return
		}
		if age < 0 || age > 100 {
			c.AbortWithStatusJSON(http.StatusBadRequest, "输入的数据错误，年龄数据错误")
			return
		}
		if len(name) < 6 || len(name) > 12 {
			c.AbortWithStatusJSON(http.StatusBadRequest, "用户名只能是6-12位")
			return
		}
		c.Next() //执行后续操作
		fmt.Println(name, age)
	}
}

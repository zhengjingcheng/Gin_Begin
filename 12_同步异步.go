package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()
	//同步请求
	r.GET("sync", func(c *gin.Context) {
		sync(c)
		c.JSON(200, ">>>主程序（主go程）同步已经执行<<<")
	})
	//异步请求
	r.GET("/async", func(c *gin.Context) {
		for i := 0; i < 6; i++ {
			cp := c.Copy()
			go async(cp, i)
		}
		c.JSON(200, "^^^主程序（主go程）异步已经执行^^^")
	})
	r.Run(":9090")
}

func async(cp *gin.Context, i int) {
	fmt.Println("第" + strconv.Itoa(i) + "个go程开始执行：" + cp.Request.URL.Path)
	time.Sleep(time.Second * 3)
	fmt.Println("第" + strconv.Itoa(i) + "个go程执行结束+++++")
}

func sync(c *gin.Context) {
	println("开始执行同步任务:" + c.Request.URL.Path)
	time.Sleep(time.Second * 3)
	println("同步任务执行完成!")
}

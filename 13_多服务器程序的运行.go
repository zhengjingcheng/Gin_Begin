package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var g errgroup.Group

func main() {
	//服务器1：
	server01 := &http.Server{
		Addr:        ":9091",
		Handler:     router01(),
		ReadTimeout: 5 * time.Second,
	}
	server02 := &http.Server{
		Addr:        ":9092",
		Handler:     router02(),
		ReadTimeout: 5 * time.Second,
	}
	g.Go(func() error { //开启服务器1
		return server01.ListenAndServe()
	})
	g.Go(func() error { //开启服务器1
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		fmt.Println("执行失败：", err)
	}
}

func router01() http.Handler {
	r1 := gin.Default()
	r1.GET("myserver1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "服务器1",
		})
	})
	return r1
}
func router02() http.Handler {
	r1 := gin.Default()
	r1.GET("myserver2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "服务器2",
		})
	})
	return r1
}

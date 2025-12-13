package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//配置路由
	r.GET("/",func(c *gin.Context)){
		c.String(200,"值:%v","你好gin")
	}
	r.Run()//启动一个web服务
}

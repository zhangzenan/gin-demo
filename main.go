package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值:%v", "你好gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})
	r.POST("/add", func(ctx *gin.Context) {
		ctx.String(200, "我是post请求返回的数据")
	})
	r.PUT("/edit", func(ctx *gin.Context) {
		ctx.String(200, "这是PUT请求")
	})
	r.DELETE("/delete", func(ctx *gin.Context) {
		ctx.String(200, "这是delete请求")
	})
	r.GET("json1", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})
	r.GET("json2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin-22",
		})
	})
	r.GET("json3", func(ctx *gin.Context) {
		a := &Article{
			Title:   "我是一个标题",
			Content: "测试内容",
			Desc:    "描述",
		}
		ctx.JSON(200, a)
	})
	r.Run(":8000") //启动一个web服务
}

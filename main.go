package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 定义路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// 启动服务
	router.Run(":8888")
}

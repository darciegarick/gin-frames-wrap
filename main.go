package main

import (
	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/bootstrap"
	"github.com/showmebug/my-gin-demo/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	router := gin.Default()
	// 定义路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// 启动服务
	router.Run(":" + global.App.Config.App.Port)
}

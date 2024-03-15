package main

import (
	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/bootstrap"
	"github.com/showmebug/my-gin-demo/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

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

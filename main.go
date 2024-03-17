/*
 * @Author: zhuziying 1249353194@qq.com
 * @Date: 2024-03-13 16:00:31
 * @LastEditors: zhuziying 1249353194@qq.com
 * @LastEditTime: 2024-03-17 18:32:04
 * @FilePath: /smb-code/smb-go/myproject/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
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

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 启动服务器
	bootstrap.RunServer()
}

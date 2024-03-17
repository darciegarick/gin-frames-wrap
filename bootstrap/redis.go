/*
 * @Author: zhuziying 1249353194@qq.com
 * @Date: 2024-03-17 14:37:46
 * @LastEditors: zhuziying 1249353194@qq.com
 * @LastEditTime: 2024-03-17 14:45:52
 * @FilePath: /smb-code/smb-go/myproject/bootstrap/redis.go
 * @Description: 初始化 Redis
 */
package bootstrap

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/showmebug/my-gin-demo/global"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Host,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}

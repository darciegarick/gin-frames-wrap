/*
 * @Author: zhuziying 1249353194@qq.com
 * @Date: 2024-03-13 16:52:28
 * @LastEditors: zhuziying 1249353194@qq.com
 * @LastEditTime: 2024-03-17 18:38:22
 * @FilePath: /smb-code/smb-go/myproject/global/app.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/showmebug/my-gin-demo/configs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)

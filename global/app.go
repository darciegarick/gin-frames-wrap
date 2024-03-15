package global

import (
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
}

var App = new(Application)

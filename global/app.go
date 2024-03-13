package global

import (
	"github.com/showmebug/my-gin-demo/configs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Configuration
	Log         *zap.Logger
}

var App = new(Application)

package global

import (
	"github.com/showmebug/my-gin-demo/configs"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      configs.Configuration
}

var App = new(Application)

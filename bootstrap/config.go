package bootstrap

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/showmebug/my-gin-demo/global"
	"github.com/spf13/viper"
)

// 初始化配置
func InitializeConfig() *viper.Viper {
	// 设置配置文件路径
	config := "config.yaml"
	// 生产环境可以通过设置环境变量来改变配置文件路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	//如果读取配置文件失败，这行代码会引发 panic 错误，并将错误信息包装为一个格式化的字符串
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return v
}

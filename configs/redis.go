/*
 * @Author: zhuziying 1249353194@qq.com
 * @Date: 2024-03-17 14:25:11
 * @LastEditors: zhuziying 1249353194@qq.com
 * @LastEditTime: 2024-03-17 14:31:53
 * @FilePath: /smb-code/smb-go/myproject/configs/redis.go
 * @Description: redis 配置文件
 */
package configs

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

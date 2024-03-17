/*
 * @Author: zhuziying 1249353194@qq.com
 * @Date: 2024-03-15 18:42:43
 * @LastEditors: Zane Zhu darciegarick@gmail.com
 * @LastEditTime: 2024-03-18 00:38:35
 * @FilePath: /smb-code/smb-go/myproject/configs/jwt.go
 * @Description: jwt 配置文件
 */
package configs

type Jwt struct {
	Secret                  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl                  int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"`                                                          // token 有效期（秒）
	JwtBlacklistGracePeriod int64  `mapstructure:"jwt_blacklist_grace_period" json:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"` // 黑名单宽限时间（秒）
	RefreshGracePeriod      int64  `mapstructure:"refresh_grace_period" json:"refresh_grace_period" yaml:"refresh_grace_period"`                   // token 自动刷新宽限时间（秒）
}

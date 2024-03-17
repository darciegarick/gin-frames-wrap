/*
 * @Author: Zane Zhu darciegarick@gmail.com
 * @Date: 2024-03-18 00:35:37
 * @LastEditors: Zane Zhu darciegarick@gmail.com
 * @LastEditTime: 2024-03-18 00:35:43
 * @FilePath: /smb-code/smb-go/myproject/internal/pkg/str.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package pkg

import (
	"math/rand"
	"time"
)

// RandString 用于生成锁标识，防止任何客户端都能解锁
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

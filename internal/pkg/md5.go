/*
 * @Author: zhuziying 1249353194@qq.com
 * @Date: 2024-03-17 18:43:09
 * @LastEditors: zhuziying 1249353194@qq.com
 * @LastEditTime: 2024-03-17 18:43:13
 * @FilePath: /smb-code/smb-go/myproject/internal/pkg/md5.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package pkg


import (
    "crypto/md5"
    "encoding/hex"
)

func MD5(str []byte, b ...byte) string {
    h := md5.New()
    h.Write(str)
    return hex.EncodeToString(h.Sum(b))
}
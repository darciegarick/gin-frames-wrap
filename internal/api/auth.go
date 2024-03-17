/*
 * @Author: zhuziying 1249353194@qq.com
 * @Date: 2024-03-15 18:56:35
 * @LastEditors: zhuziying 1249353194@qq.com
 * @LastEditTime: 2024-03-17 18:51:00
 * @FilePath: /smb-code/smb-go/myproject/internal/api/auth.go
 * @Description: api 接口文件
 */
package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/showmebug/my-gin-demo/internal/common/request"
	"github.com/showmebug/my-gin-demo/internal/common/response"
	"github.com/showmebug/my-gin-demo/internal/services"
)

// 登录
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// 详情
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

// 退出
func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}

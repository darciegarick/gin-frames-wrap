package services

import (
	"errors"

	"github.com/showmebug/my-gin-demo/global"
	"github.com/showmebug/my-gin-demo/internal/common/request"
	"github.com/showmebug/my-gin-demo/internal/pkg"
	"github.com/showmebug/my-gin-demo/internal/repository"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user repository.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&repository.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = repository.User{Name: params.Name, Mobile: params.Mobile, Password: pkg.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

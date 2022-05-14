package service

import (
	"zshf.private/global"
	"zshf.private/models"
)

type WebUserService struct {
}

func (userService *WebUserService) Login(param models.WebUserLoginParam) uint64 {
	var user models.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	return user.Id
}

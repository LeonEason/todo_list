package service

import (
	"github.com/jinzhu/gorm"
	"todo/model"
	"todo/pkg/utils"
	"todo/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=16"`
}

func (service *UserService) Register() *serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return &serializer.Response{
			Status: 400,
			Msg:    "注册重复",
		}
	}
	user.UserName = service.UserName
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "数据库操作失误",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	}
}
func (service *UserService) Login() *serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return &serializer.Response{
				Status: 400,
				Msg:    "用户不存在",
			}
		}
		return &serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return &serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return &serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}

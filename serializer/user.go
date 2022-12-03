package serializer

import "todo/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	UserName string `json:"user_name" form:"user_name" example:"FanOne"`
	Status   string `json:"status" form:"status"`       //用户状态
	CreateAt int64  `json:"create_at" form:"create_at"` //创建
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}

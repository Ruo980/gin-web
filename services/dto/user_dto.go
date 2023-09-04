// Package dto
// @Description: 用户登录表映射实体user_dto
package dto

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,email" message:"用户名填写错误"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

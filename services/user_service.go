// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 15:45
// @Description:
package services

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/services/dto"
	"errors"
)

var userService *UserService

type UserService struct {
	BaseService
	userDAO *dao.UserDAO
}

func NewUserService() *UserService {
	if userService == nil { // 单例模式
		userService = &UserService{
			userDAO: dao.NewUserDAO(),
		}
	}
	return userService
}

// 定义用户登录业务：根据用户名和密码判断账号密码是否正确，正确则返回用户信息
func (s UserService) Login(userLoginDto dto.UserLoginDTO) (model.User, error) {
	var errorResult error
	user := s.userDAO.GetUserByNameAndPassword(userLoginDto.Name, userLoginDto.Password)
	if user.ID == 0 {
		errorResult = errors.New("Invalid user name or password")
	}
	return user, errorResult

}

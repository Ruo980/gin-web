// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 15:45
// @Description:
package services

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO}
}

func (service *UserService) CreateUser(user *model.User) error {
	return service.userDAO.Create(user)
}

func (service *UserService) GetUserByID(id uint) (*model.User, error) {
	return service.userDAO.FindByID(id)
}

func (service *UserService) GetAllUsers() ([]model.User, error) {
	return service.userDAO.FindAll()
}

func (service *UserService) DeleteUser(user *model.User) error {
	return service.userDAO.Delete(user)
}

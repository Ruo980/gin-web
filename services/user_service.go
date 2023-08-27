// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 15:45
// @Description:
package services

import (
	"awesomeProject/dao"
	"awesomeProject/models"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO}
}

func (service *UserService) CreateUser(user *models.User) error {
	return service.userDAO.Create(user)
}

func (service *UserService) GetUserByID(id uint) (*models.User, error) {
	return service.userDAO.FindByID(id)
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.userDAO.FindAll()
}

func (service *UserService) DeleteUser(user *models.User) error {
	return service.userDAO.Delete(user)
}

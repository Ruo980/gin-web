// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:07
// @Description:
package dao

import (
	"awesomeProject/models"
	"github.com/jinzhu/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db}
}

func (dao *UserDAO) Create(user *models.User) error {
	return dao.db.Create(user).Error
}

func (dao *UserDAO) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := dao.db.First(&user, id).Error
	return &user, err
}

func (dao *UserDAO) FindAll() ([]models.User, error) {
	var users []models.User
	err := dao.db.Find(&users).Error
	return users, err
}

func (dao *UserDAO) Delete(user *models.User) error {
	return dao.db.Delete(user).Error
}

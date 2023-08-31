// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:07
// @Description:
package dao

import (
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db}
}

func (dao *UserDAO) Create(user *model.User) error {
	return dao.db.Create(user).Error
}

func (dao *UserDAO) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := dao.db.First(&user, id).Error
	return &user, err
}

func (dao *UserDAO) FindAll() ([]model.User, error) {
	var users []model.User
	err := dao.db.Find(&users).Error
	return users, err
}

func (dao *UserDAO) Delete(user *model.User) error {
	return dao.db.Delete(user).Error
}

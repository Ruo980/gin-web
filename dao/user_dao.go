// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:07
// @Description:
package dao

import "awesomeProject/model"

var userDao *UserDAO

type UserDAO struct {
	BaseDao
}

func NewUserDAO() *UserDAO {
	if userDao == nil { // 单例模式
		userDao = &UserDAO{BaseDao: *NewBaseDao()}
	}
	return userDao
}

// 定义查询用户的方法：根据用户名和密码查询用户
func (m UserDAO) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var user model.User
	m.Orm.Where("user_name = ? and password = ?", stUserName, stPassword).First(&user)
	return user
}

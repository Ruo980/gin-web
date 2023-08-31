// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:05
// @Description:
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null" json:"name"`
	Avatar   string `gorm:"size:255" json:""`
	Mobile   string `gorm:"size:11"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"size:128;not null" json:"-"`
}

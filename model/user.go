// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:05
// @Description:
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // gorm.Model 包含了ID、CreatedAt、UpdatedAt、DeletedAt四个字段
	Name       string `gorm:"size:64;not null" json:"name"` // 用户名、账号：保持对话的标识
	Avatar     string `gorm:"size:255" json:""`
	Mobile     string `gorm:"size:11"`
	Email      string `gorm:"size:128"`
	Password   string `gorm:"size:128;not null" json:"-"`
}

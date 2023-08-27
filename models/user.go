// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:05
// @Description:
package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Email     string     `gorm:"not null;unique" json:"email"`
	Password  string     `gorm:"not null" json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

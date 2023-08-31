// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:05
// @Description:
package model

import "time"

type Product struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Price     float64    `gorm:"not null" json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

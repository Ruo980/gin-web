// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:07
// @Description:
package dao

import (
	"awesomeProject/models"
	"github.com/jinzhu/gorm"
)

type ProductDAO struct {
	db *gorm.DB
}

func NewProductDAO(db *gorm.DB) *ProductDAO {
	return &ProductDAO{db}
}

func (dao *ProductDAO) Create(product *models.Product) error {
	return dao.db.Create(product).Error
}

func (dao *ProductDAO) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := dao.db.First(&product, id).Error
	return &product, err
}

func (dao *ProductDAO) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := dao.db.Find(&products).Error
	return products, err
}

func (dao *ProductDAO) Delete(product *models.Product) error {
	return dao.db.Delete(product).Error
}

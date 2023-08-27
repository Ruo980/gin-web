// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 15:45
// @Description:
package services

import (
	"awesomeProject/dao"
	"awesomeProject/models"
)

type ProductService struct {
	productDAO *dao.ProductDAO
}

func NewProductService(productDAO *dao.ProductDAO) *ProductService {
	return &ProductService{productDAO}
}

func (service *ProductService) CreateProduct(product *models.Product) error {
	return service.productDAO.Create(product)
}

func (service *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return service.productDAO.FindByID(id)
}

func (service *ProductService) GetAllProducts() ([]models.Product, error) {
	return service.productDAO.FindAll()
}

func (service *ProductService) DeleteProduct(product *models.Product) error {
	return service.productDAO.Delete(product)
}

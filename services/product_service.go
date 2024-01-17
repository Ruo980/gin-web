// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 15:45
// @Description:
package services

import (
	"awesomeProject/dao"
	"awesomeProject/model"
)

var productService *ProductService

type ProductService struct {
	BaseService
	productDao *dao.ProductDao
}

// NewProductService
//
//	@Description: 单例模式
//	@return *ProductService
func NewProductService() *ProductService {
	if productService == nil { // 单例模式
		productService = &ProductService{
			productDao: dao.NewProductDao(),
		}
	}
	return productService
}

// GetProductByName
//
//	@Description: 根据商品名称查询商品
//	@receiver s
//	@param stProductName
//	@return model.Product
func (s ProductService) GetProductByName(stProductName string) model.Product {
	return s.productDao.GetProductByName(stProductName)
}

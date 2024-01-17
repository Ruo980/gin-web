package dao

import "awesomeProject/model"

var productDao *ProductDao

type ProductDao struct {
	BaseDao
}

// NewProductDao
//
//	@Description: 构造函数单例模式，返回对象指针
//	@return *ProductDao
func NewProductDao() *ProductDao {
	if productDao == nil {
		productDao = &ProductDao{BaseDao: *NewBaseDao()}
	}
	return productDao
}

// GetProductByName
//
//	@Description: 根据商品名称查询商品
//	@receiver m
//	@param stProductName
//	@return model.Product
func (m ProductDao) GetProductByName(stProductName string) model.Product {
	var product model.Product
	m.Orm.Where("product_name = ?", stProductName).First(&product)
	return product
}

// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-22 15:49
// @Description:
package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct{}

func (ProductController) ListProductHandler(c *gin.Context) {
	// 获取产品列表的处理逻辑
	fmt.Println("获取列表")
	c.JSON(http.StatusOK, gin.H{"messgae": "获取列表成功"})
}

func (ProductController) GetProductHandler(c *gin.Context) {
	// 根据产品id获取产品信息的处理逻辑
	fmt.Println("获取指定产品信息")
	c.JSON(http.StatusOK, gin.H{"messgae": "获取指定产品信息成功"})
}

func (ProductController) CreateProductHandler(c *gin.Context) {
	// 创建产品的处理逻辑
	fmt.Println("创建一个产品信息")
	c.JSON(http.StatusOK, gin.H{"messgae": "创建一个产品信息成功"})
}

// 添加其他产品相关路由处理函数的实现

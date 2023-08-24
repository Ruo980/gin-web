// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-22 15:42
// @Description:
package routers

import (
	"awesomeProject/controllers/product"
	"github.com/gin-gonic/gin"
)

func SetupProductRouterInit(router *gin.Engine) {
	productRouter := router.Group("/product")
	{ //这个{}没有实际意义，只是方便阅读
		productRouter.GET("", product.ProductController{}.ListProductHandler)
		productRouter.GET("/:id", product.ProductController{}.GetProductHandler)
		productRouter.POST("", product.ProductController{}.CreateProductHandler)
		// 添加其他产品相关路由处理函数
	}
}

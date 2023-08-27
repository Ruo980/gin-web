// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:16
// @Description:
package routers

import (
	"awesomeProject/handlers"
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
)

func SetupProductRouter(productService *services.ProductService) *gin.Engine {
	router := gin.Default()

	productHandler := handlers.NewProductHandler(productService)

	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.GET("/products", productHandler.GetAllProducts)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	return router
}

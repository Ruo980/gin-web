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

func SetupUserRouter(userService *services.UserService) *gin.Engine {
	router := gin.Default()

	userHandler := handlers.NewUserHandler(userService)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUserByID)
	router.GET("/users", userHandler.GetAllUsers)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	return router
}

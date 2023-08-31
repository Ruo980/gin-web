// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 14:16
// @Description:
package routers

import (
	"awesomeProject/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRouter() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		rgPublicUser.POST("login", userApi.Login)
		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "zs"},
					{"id": 2, "name": "ls"},
				},
			})
		})
		rgAuthUser.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "zs",
			})
		})
	})
}

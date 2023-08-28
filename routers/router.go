// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-28 11:14
// @Description: 路由初始化文件
package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 定义一个函数形式结构体
type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegistRoute
)

func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitRouter() {
	r := gin.Default()

	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	// 路由注册：调用数组中的函数实现路由注册
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	//读取配置文件服务端口，设置run
	stPort := viper.Get("server.port")
	if stPort == "" {
		stPort = "8882"
	}
	err := r.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("Start Server Error:%s", err.Error()))
	}
}

func InitBasePlatformRoutes() {
	// 初始化user模块路由
	InitUserRouter()
}

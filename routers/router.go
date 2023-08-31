// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-28 11:14
// @Description: 路由初始化文件
package routers

import (
	_ "awesomeProject/docs"
	"awesomeProject/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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

	// 启动时创建可取消的上下文并监听：输入ctrl+c会触发应用退出该上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// 初始化Gin框架并注册路由
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	// 路由注册：调用数组中的函数实现路由注册
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 读取配置文件服务端口，设置run
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8882"
	}

	// 启动服务
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		// 启动成功
		global.Logger.Info(fmt.Sprintf("Start Server Listen:%s", stPort))
		// 启动失败处理
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Server Error:%s", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("stop Server Error:%s", err.Error()))
		return
	}
	global.Logger.Info("stop Server Success")

}

func InitBasePlatformRoutes() {
	// 初始化user模块路由
	InitUserRouter()
}

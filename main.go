// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-22 15:39
// @Description: 应用程序入口函数
package main

import (
	"awesomeProject/dao"
	"awesomeProject/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// 注册路由
	routers.SetupUserRouter(router)
	routers.SetupProductRouter(router)

	// 初始化数据库连接
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	router.Run(":9999")

}

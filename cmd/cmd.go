// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-28 10:29
// @Description:
package cmd

import (
	"awesomeProject/conf"
	"awesomeProject/global"
	"awesomeProject/routers"
	"fmt"
)

func Start() {
	// 初始化系统配置
	conf.InitConfig()
	global.Logger = conf.InitLogger()
	routers.InitRouter()
}

func Clean() {
	fmt.Println("=======Clean=======")
}

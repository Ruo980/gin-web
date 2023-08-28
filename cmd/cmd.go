// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-28 10:29
// @Description:
package cmd

import (
	"awesomeProject/conf"
	"awesomeProject/routers"
	"fmt"
)

func Start() {
	conf.InitConfig()
	routers.InitRouter()
}

func Clean() {
	fmt.Println("=======Clean=======")
}

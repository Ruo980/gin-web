// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-22 15:39
// @Description:
package main

import (
	"awesomeProject/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routers.SetupProductRouterInit(router)

	router.Run(":9999")

}

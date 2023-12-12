// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-28 10:29
// @Description:
package cmd

import (
	"awesomeProject/conf"
	"awesomeProject/global"
	"awesomeProject/routers"
	"awesomeProject/utils"
	"fmt"
)

func Start() {
	//定义一个初始化错误，存放初始化过程中的所有错误
	var initErr error
	// 初始化系统配置
	conf.InitConfig()
	// 初始化日志系统
	global.Logger = conf.InitLogger()
	// 初始化数据库连接
	db, err := conf.InitDB()
	if err != nil {
		// 将数据库初始化错误保存起来，归纳到初始化错误集中
		utils.AppendError(initErr, err)
	}
	global.DB = db
	// 初始化Redis连接
	global.RedisClient, err = conf.InitRedis()
	if err != nil {
		// 将数据库初始化错误保存起来，归纳到初始化错误集中
		utils.AppendError(initErr, err)
	}

	// 判断初始化过程是否出错
	if initErr != nil {
		if global.Logger != nil {
			// 日志初始化正确，使用日志系统来记录错误
			global.Logger.Error(initErr.Error())
		}
		// 日志初始化错误，直接输出到终端
		panic(initErr.Error())
	}
	routers.InitRouter()
}

func Clean() {
	fmt.Println("=======Clean=======")
}

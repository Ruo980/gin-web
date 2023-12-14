// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-22 15:39
// @Description: 应用程序入口函数
package main

import "awesomeProject/cmd"

func main() {

	defer cmd.Clean()
	cmd.Start()

}

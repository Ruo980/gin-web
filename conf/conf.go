// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-28 10:14
// @Description: 该文件提供一个函数用于将yml配置信息读取为数据配置到Gin框架中
package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// 初始化配置函数
func InitConfig() {
	// 设置配置文件名为 "settings"
	viper.SetConfigName("settings")
	// 设置配置文件类型为 "yml"
	viper.SetConfigType("yml")
	// 找到同级目录下的settings.yml文件:'./'表示当前项目下
	viper.AddConfigPath("./conf/")
	// 读取并解析配置文件
	err := viper.ReadInConfig()

	// 如果读取配置文件出错，则抛出异常
	if err != nil {
		panic(fmt.Sprintf("Load Config Error：%s", err.Error()))
	}

	// 打印配置文件中的 "server.port" 的值
	fmt.Println(viper.GetString("server.port"))
}

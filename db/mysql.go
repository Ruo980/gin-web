// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-26 21:21
// @Description: db包存放数据库连接。对外暴露全局变量DB *gorm.DB
package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"post"`
}

// 给结构体定义一个方法:利用viper读取config/database.yaml的配置并将值赋给结构体实例
func (c *Conf) getConf() (*Conf, error) {

	viper.SetConfigName("database")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read conf file: %w", err)
	}
	var conf Conf
	err = viper.Unmarshal(&conf) //将
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal conf: %w", err)
	}

	return &conf, nil
}

var (
	DB *gorm.DB //将“连接”设置为全局变量，实现初始化和关闭函数的分离
)

// 初始化函数
func InitMySQL() (err error) {
	var c Conf
	conf, _ := c.getConf()
	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	// 连接数据库
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 验证数据库连接是否成功，若成功，则无异常
	return DB.DB().Ping()
}

// 关闭连接函数
func Close() {
	DB.Close()
}

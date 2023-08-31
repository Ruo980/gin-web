// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-31 19:00
// @Description: 数据库驱动配置
package conf

import (
	"awesomeProject/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {
	// 指定sql日志输出模式：确定哪些等级的日志被记录
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
	// 从setting.yml中获取信息，拼接dns
	// 获取数据库连接信息
	dbConfig := viper.GetStringMapString("database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["username"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["dbname"],
	)
	// 使用gorm打开mysql数据库连接:配置数据库表命名规范以及日志输出等级
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", // 设置表名前缀为sys_
			SingularTable: true,   // 禁用表名的复数形式
		},
		Logger: logger.Default.LogMode(logMode),
	})

	// 链接失败:返回链接nil
	if err != nil {
		return nil, err
	}

	// 链接成功:配置链接设置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("database.maxIdleConn")) //最大空闲连接数
	sqlDB.SetMaxOpenConns(viper.GetInt("database.maxOpenConn")) //同时打开的连接数的最大限制
	sqlDB.SetConnMaxLifetime(time.Hour)                         //连接被占用的最大生命周期时间
	db.AutoMigrate(&model.User{})
	return db, err
}

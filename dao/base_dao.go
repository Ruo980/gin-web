package dao

import (
	"awesomeProject/global"
	"gorm.io/gorm"
)

// 基础 Dao 包：将 orm 全局变量引入，避免多个dao文件中重复引入

type BaseDao struct {
	Orm *gorm.DB
}

func NewBaseDao() *BaseDao {
	return &BaseDao{
		Orm: global.DB, // 从全局变量中获取 orm 实例
	}
}

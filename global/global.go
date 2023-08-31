// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-31 13:13
// @Description:
package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.SugaredLogger
	DB     *gorm.DB
)

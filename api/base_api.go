// Package api
// @Description: 基础处理函数
package api

import (
	"awesomeProject/global"
	"awesomeProject/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

// NewBaseApi
//
//	@Description: BaseApi构造函数
//	@return BaseApi
func NewBaseApi() BaseApi {
	return BaseApi{
		// 初始化实例时带上日志
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error

	// TODO 绑定请求上下文
	m.Ctx = option.Ctx

	// 绑定请求数据
	if option.DTO != nil {
		if option.BindParamsFromUri {
			errResult = m.Ctx.ShouldBind(option.DTO)
		} else {
			errResult = m.Ctx.ShouldBind(option.DTO)
		}
		if errResult != nil {
			errResult = m.parseValidateErrors(errResult.(validator.ValidationErrors), option.DTO)
			m.AddError(errResult)
			m.Fail(ResponseJson{
				Msg: m.GetError().Error(),
			})
		}
	}
	return m
}

// AddError
//
//	@Description: 使用utils工具将处理函数中出现的错误放置到BaseApi中
//	@receiver m
//	@param errNew
func (m *BaseApi) AddError(errNew error) {
	m.Errors = utils.AppendError(m.Errors, errNew)
}

// GetError
//
//	@Description: 获取所有处理函数中的错误
//	@receiver m
//	@return error
func (m *BaseApi) GetError() error {
	return m.Errors
}

func (m *BaseApi) parseValidateErrors(err validator.ValidationErrors, target any) error {
	var errResult error
	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range err {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}
		if errMessage != "" {
			errMessage = fmt.Sprintf("%s:%s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}

func (m *BaseApi) Fail(resp ResponseJson) {
	Fail(m.Ctx, resp)

}

func (m *BaseApi) OK(resp ResponseJson) {
	OK(m.Ctx, resp)
}

// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-30 17:21
// @Description: 用户路由处理函数
package api

import (
	"awesomeProject/services/dto"
	"awesomeProject/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Tag 用户登录
// @Summary 用户登录
// @Description 用户登录详情描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string “登录成功”
// @Failure 401 {string} string “登录失败”
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(context *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	err := context.ShouldBind(&iUserLoginDTO)
	if err != nil {
		Fail(context, ResponseJson{
			Msg: parseValidateErrors(err.(validator.ValidationErrors), &iUserLoginDTO).Error(),
		})
		return
	}
	OK(context, ResponseJson{
		Msg:  "Login Success",
		Data: iUserLoginDTO,
	})

	/*	Fail(context, ResponseJson{
		Msg: "Login Failed",
	})*/
}

func parseValidateErrors(err validator.ValidationErrors, target any) error {
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
		if errMessage == "" {
			errMessage = fmt.Sprintf("%s:%s Error", fieldErr.Field(), fieldErr.Tag())
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}

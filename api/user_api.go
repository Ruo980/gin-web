// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-30 17:21
// @Description: 用户路由处理函数
package api

import (
	"awesomeProject/services/dto"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
	}
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
	/*	err := context.ShouldBind(&iUserLoginDTO)
		if err != nil {
			Fail(context, ResponseJson{
				Msg: parseValidateErrors(err.(validator.ValidationErrors), &iUserLoginDTO).Error(),
			})
			return
		}*/
	if err := m.BuildRequest(BuildRequestOption{Ctx: context, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	m.OK(ResponseJson{
		Msg:  "Login Success",
		Data: iUserLoginDTO,
	})
	/*OK(context, ResponseJson{
		Msg:  "Login Success",
		Data: iUserLoginDTO,
	})*/

	/*	Fail(context, ResponseJson{
		Msg: "Login Failed",
	})*/
}

// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-30 17:21
// @Description: 用户路由处理函数
package api

import (
	"awesomeProject/services"
	"awesomeProject/services/dto"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	service *services.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		service: services.NewUserService(),
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
	/*
			将以下部分的操作提取到BaseApi中，实现解析操作的分离
		err := context.ShouldBind(&iUserLoginDTO)
			if err != nil {
				Fail(context, ResponseJson{
					Msg: parseValidateErrors(err.(validator.ValidationErrors), &iUserLoginDTO).Error(),
				})
				return
			}*/
	if err := m.BuildRequest(BuildRequestOption{Ctx: context, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	// 调用 service 层的 Login 方法，传入 iUserLoginDTO ，返回 user 和 error
	user, err := m.service.Login(iUserLoginDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	// 生成 token ，并将 token 响应给前端
	token, err := utils.GenerateToken(user.ID, user.Name)
	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}

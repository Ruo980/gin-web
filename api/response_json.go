// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-09-01 16:31
// @Description: 结果响应类
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	// HTTP 响应状态码
	Status int `json:"-"` // 状态码有 HTTP response 自带
	// 业务响应状态码
	Code int `json:"code,omitempty"`
	// 业务响应消息
	Msg string `json:"msg,omitempty"`
	// 响应数据
	Data any `json:"data,omitempty"`
}

// 判断响应结果是否为空值
func (m ResponseJson) isEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

// 响应方法：响应前端请求，需要提供 http 状态码和 ResponseJson 两个数据。
// 借助 *gin.Context 首先传入状态和响应的 responseJson ，然后判断responseJson是否为空，最好根据状态和responseJson返回。
func HttpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.isEmpty() {
		// 只返回一个
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

// 手动构建状态Status
func buildStatus(resp ResponseJson, nDefaultStatus int) int {
	if 0 == resp.Status {
		return nDefaultStatus
	}
	return resp.Status
}

func OK(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

func Fail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

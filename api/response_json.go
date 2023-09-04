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
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

// 判断响应结果是否为空值
func (m ResponseJson) isEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

// 响应方法：响应前端请求：需要提供http状态码和ResponseJson两个数据
// 首先传入状态和响应的responseJson，然后判断responseJson是否为空，最好根据状态和responseJson返回。
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

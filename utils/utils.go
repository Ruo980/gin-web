// @Program: awesomeProject
// @Author: FunCoder
// @Create: 2023-08-31 22:47
// @Description: 错误集合
package utils

import "fmt"

// 错误存放工具
func AppendError(existErr, newError error) error {
	if existErr == nil {
		existErr = newError
	}
	return fmt.Errorf("%v,%w", existErr, newError)
}

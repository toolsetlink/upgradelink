// Package osx 获取系统环境变量
//
//	@Desc: 获取系统环境变量
package osx

import (
	"os"

	"github.com/zeromicro/go-zero/core/service"
)

var env = service.DevMode

func init() {
	if e, ok := os.LookupEnv("Env"); ok {
		env = e
	}
}

// Env 获取当前环境变量
//
//	@desc    获取当前环境变量
//	@param
//	@return
func Env() string {
	return env
}

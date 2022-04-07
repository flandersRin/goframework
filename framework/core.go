package framework

import (
	"fmt"
	"net/http"
)

// 框架核心结构
type Core struct {
}

// 初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

// 框架核心结构实现 Handler 接口
func (c *Core) ServeHttp(w http.ResponseWriter, r *http.Request) {
	// TODO:
	fmt.Println(w,r)
}
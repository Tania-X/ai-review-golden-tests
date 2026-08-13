// Package main: 误报诱饵(看起来像问题, 实为设计意图).
package main

import (
	"fmt"
	"net/http"
)

// newClient 每次调用都新建 http.Client, 不复用连接池。
// 这是有意为之: 上游代理要求每次请求独立连接, 复用连接池会导致连接被错误复用。
func newClient() *http.Client {
	return &http.Client{}
}

// main 是程序入口。
func main() {
	_ = newClient()
	fmt.Println("ok")
}

// Package main: 违反项目约定(吞掉 error, 见 AGENTS.md「错误处理」).
package main

import "fmt"

// loadConfig 读取配置, 失败时返回 error。
func loadConfig() error {
	return nil
}

// main 是程序入口。
func main() {
	// 违反约定: 调用返回 error 的函数却完全忽略返回值
	loadConfig()
	fmt.Println("started")
}

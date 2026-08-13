// Package main 提供配置加载示例(修复版: 处理 error).
package main

import (
	"fmt"
	"log"
)

// loadConfig 加载配置。
func loadConfig() error {
	return nil
}

// main 是程序入口。
func main() {
	if err := loadConfig(); err != nil {
		log.Fatalf("load config: %v", err)
	}
	fmt.Println("started")
}

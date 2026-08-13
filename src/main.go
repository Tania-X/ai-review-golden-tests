// Package main 提供配置加载示例(修复版: 处理 error).
package main

import (
	"fmt"
	"log"
	"os"
)

// loadConfig 从配置文件读取配置。
func loadConfig() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	_ = data // 简化: 本例不解析内容
	return nil
}

// main 是程序入口。
func main() {
	if err := loadConfig(); err != nil {
		log.Fatalf("load config: %v", err)
	}
	fmt.Println("started")
}

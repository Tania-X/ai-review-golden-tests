// Package main 提供求和示例.
package main

import "fmt"

// User 表示一个用户。
type User struct {
	ID   int
	Name string
}

// sum 计算整数切片之和。
func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// main 是程序入口。
func main() {
	nums := []int{1, 2, 3, 4, 5}
	_, _ = fmt.Printf("sum=%d\n", sum(nums)) // 显式接收返回值(输出到 stdout 失败可忽略)
}

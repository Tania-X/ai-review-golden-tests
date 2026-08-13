// Package main: 注入 nil 指针解引用 bug(相比基线).
package main

import "fmt"

// User 表示一个用户。
type User struct {
	ID   int
	Name string
}

var cache = map[string]*User{}

// getUser 从缓存取用户, 可能返回 nil。
func getUser(id string) *User {
	return cache[id]
}

// getName 返回用户名; 未考虑 getUser 返回 nil 的情况。
func getName(id string) string {
	u := getUser(id)
	return u.Name // nil 解引用: id 不存在时 panic
}

func main() {
	fmt.Println(getName("missing"))
}

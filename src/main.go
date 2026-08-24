// Package main 提供用户查询示例.
package main

import "fmt"

// User 表示一个用户。
type User struct {
	ID   int
	Name string
}

var cache = map[string]*User{}

// getUser 根据 id 返回用户。
func getUser(id string) *User {
	return cache[id]
}

// getName 返回用户姓名。
func getName(id string) string {
	return getUser(id).Name
}

func main() {
	fmt.Println(getName("missing"))
}

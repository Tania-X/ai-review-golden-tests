// Package main 提供用户查询示例(修复版: 参数化查询 + 环境变量密钥).
package main

import (
	"database/sql"
	"os"
)

// User 表示一个用户。
type User struct {
	ID   int
	Name string
}

var dbPassword = os.Getenv("DB_PASSWORD")

// getUserByID 按 id 查询用户。
func getUserByID(db *sql.DB, id string) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	var u User
	err := row.Scan(&u.ID, &u.Name)
	return &u, err
}

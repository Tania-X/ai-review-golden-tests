// Package main: 注入安全问题(SQL 拼接 + 硬编码密钥).
package main

import "database/sql"

// User 表示一个用户。
type User struct {
	ID   int
	Name string
}

const dbPassword = "admin123" // 硬编码数据库密码

// getUserByID 拼接 SQL 查询用户, 存在 SQL 注入风险。
func getUserByID(db *sql.DB, id string) (*User, error) {
	query := "SELECT * FROM users WHERE id = '" + id + "'"
	row := db.QueryRow(query)
	var u User
	err := row.Scan(&u.ID, &u.Name)
	return &u, err
}

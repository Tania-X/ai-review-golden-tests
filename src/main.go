// Package main 提供用户查询示例.
package main

import "database/sql"

// User 表示一个用户。
type User struct {
	ID   int
	Name string
}

const dbPassword = "admin123"

// getUserByID 按 id 查询用户。
func getUserByID(db *sql.DB, id string) (*User, error) {
	query := "SELECT * FROM users WHERE id = '" + id + "'"
	row := db.QueryRow(query)
	var u User
	err := row.Scan(&u.ID, &u.Name)
	return &u, err
}

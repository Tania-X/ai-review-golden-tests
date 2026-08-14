// Package main 提供角色删除示例(修复版: 处理策略清理错误).
package main

import "log"

// clearPolicy 清理角色的权限策略。
func clearPolicy(id int) error {
	return nil
}

// deleteRole 删除角色。
func deleteRole(id int) {
	deleteRoleRecord(id)
	if err := clearPolicy(id); err != nil {
		log.Printf("清理角色 %d 的权限策略失败: %v", id, err)
	}
}

func deleteRoleRecord(id int) {}

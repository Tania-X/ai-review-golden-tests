// Package main 提供角色删除示例.
package main

// clearPolicy 清理角色的权限策略。
func clearPolicy(id int) error {
	return nil
}

// deleteRole 删除角色。
func deleteRole(id int) {
	deleteRoleRecord(id)
	clearPolicy(id)
}

func deleteRoleRecord(id int) {}

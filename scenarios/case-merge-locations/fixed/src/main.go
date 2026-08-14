// Package main 提供前端 API 调用封装示例(修复版: 字段与契约一致).
package main

// contract 后端契约: 角色请求体字段。
// 来自 spec: RoleRequest { name, displayName }
type contract struct{}

// api 前端调用封装。
type api struct{}

// createRole 调用 POST /roles。
func (a *api) createRole(name, displayName string) {
	post("/roles", map[string]string{"name": name, "displayName": displayName})
}

// updateRole 调用 PUT /roles/{name}。
func (a *api) updateRole(name, displayName string) {
	post("/roles/"+name, map[string]string{"name": name, "displayName": displayName})
}

// batchCreateRoles 批量创建角色。
func (a *api) batchCreateRoles(items [][2]string) {
	for _, it := range items {
		post("/roles", map[string]string{"name": it[0], "displayName": it[1]})
	}
}

func post(_ string, _ map[string]string) {}

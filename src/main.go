// Package main 提供前端 API 调用封装示例(与后端契约对照).
package main

// contract 后端契约: 角色请求体字段。
// 来自 spec: RoleRequest { name, displayName }
type contract struct{}

// api 前端调用封装。
type api struct{}

// createRole 调用 POST /roles。
func (a *api) createRole(name, label string) {
	post("/roles", map[string]string{"name": name, "label": label})
}

// updateRole 调用 POST /roles/{name}。
func (a *api) updateRole(name, label string) {
	post("/roles/"+name, map[string]string{"name": name, "label": label})
}

// batchCreateRoles 批量创建角色。
func (a *api) batchCreateRoles(items [][2]string) {
	for _, it := range items {
		post("/roles", map[string]string{"name": it[0], "label": it[1]})
	}
}

func post(_ string, _ map[string]string) {}

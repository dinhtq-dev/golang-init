package constants

// PermissionGroup đại diện cho một nhóm quyền
type PermissionGroup struct {
	GroupName   string   // Tên của nhóm quyền
	Permissions []string // Danh sách quyền thuộc nhóm
}

// Danh sách quyền được cấu hình sẵn
var PermissionGroups = []PermissionGroup{
	{
		GroupName:   "user",
		Permissions: []string{"user_create", "user_read", "user_update"},
	},
	{
		GroupName:   "admin",
		Permissions: []string{"admin_manage_users", "admin_manage_roles", "admin_access_dashboard"},
	},
	{
		GroupName:   "product",
		Permissions: []string{"product_create", "product_read", "product_update", "product_delete"},
	},
}

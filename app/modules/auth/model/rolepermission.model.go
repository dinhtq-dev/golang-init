package model
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
	AssignedAt   int64 `json:"assigned_at"`
	AssignedBy   uint  `json:"assigned_by"`
}
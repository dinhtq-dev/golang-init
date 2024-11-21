package model

type Role struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}
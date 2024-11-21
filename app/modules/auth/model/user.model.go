package model

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      *string   `gorm:"type:varchar(255)" json:"name"`
    Email     string    `gorm:"type:varchar(255);unique" json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Roles []Role `json:"roles" gorm:"many2many:user_roles;"`
}

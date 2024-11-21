package seeder

import (
	"golang/app/modules/auth/model"
	"log"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []model.User{
		{Name: nil, Email: "0o7YV@example.com", Roles: []model.Role{{Name: "admin"}}},
		{Name: nil, Email: "3243@example.com", Roles: []model.Role{{Name: "admin"}}},
		{Name: nil, Email: "4324234@example.com", Roles: []model.Role{{Name: "admin"}}},
	}

	for _, user := range users {
		err := db.Where(model.User{Email: user.Email}).FirstOrCreate(&user).Error
		if err != nil {
			log.Printf("Failed to seed user %s: %v", user.Email, err)
		} else {
			log.Printf("Seeded user: %s", user.Email)
		}
	}
}

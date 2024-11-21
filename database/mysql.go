// app/database/mysql.go
package database

import (
	"fmt"
	"golang/app/modules/auth/model"
	"golang/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB initializes the MySQL database connection
func InitDB(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	
	// Connect to MySQL database
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // Data Source Name (DSN)
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Enable debug mode to log SQL queries
	// db = db.Debug()

	// Automatically migrate the User model to create the corresponding table
	err = db.AutoMigrate(&model.User{}, &model.UserRole{}, &model.Role{}, &model.Permission{}, &model.RolePermission{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	return db
}

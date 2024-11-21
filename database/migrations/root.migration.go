package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"time"

	"gorm.io/gorm"
)

// Migration chỉ lưu thông tin về trạng thái migration
type Migration struct {
	ID           uint       `gorm:"primaryKey"`
	Name         string     `gorm:"unique;not null"`
	AppliedAt    *time.Time `gorm:"not null"`
	RolledBackAt *time.Time `gorm:"index"`
	Up   func(*gorm.DB) error
	Down func(*gorm.DB) error
}

// MigrationLogic lưu trữ logic Up và Down cho migration
type MigrationLogic struct {
	Name string
	Up   func(*gorm.DB) error
	Down func(*gorm.DB) error
}

// CreateMigrationsTable tạo bảng `migrations` nếu chưa tồn tại
func CreateMigrationsTable(db *gorm.DB) error {
	return db.AutoMigrate(&Migration{})
}

// LoadMigrations nạp các file plugin và trả về danh sách logic migration
func LoadMigrations(directory string) ([]MigrationLogic, error) {
	var migrationLogics []MigrationLogic

	// Duyệt qua tất cả file .so trong thư mục
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %s: %w", path, err)
		}

		// Bỏ qua nếu không phải file .so
		if filepath.Ext(path) != ".so" {
			return nil
		}

		// Nạp plugin
		plug, err := plugin.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open plugin %s: %w", path, err)
		}

		// Tìm hàm Up
		upFunc, err := plug.Lookup("Up")
		if err != nil {
			return fmt.Errorf("plugin %s must have an Up function: %w", path, err)
		}

		// Tìm hàm Down
		downFunc, err := plug.Lookup("Down")
		if err != nil {
			return fmt.Errorf("plugin %s must have a Down function: %w", path, err)
		}

		// Thêm logic vào danh sách
		migrationLogics = append(migrationLogics, MigrationLogic{
			Name: filepath.Base(path),
			Up:   upFunc.(func(*gorm.DB) error),
			Down: downFunc.(func(*gorm.DB) error),
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return migrationLogics, nil
}

func EnsureMigrationsTableExists(db *gorm.DB) error {
    query := `
        CREATE TABLE IF NOT EXISTS migrations (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `
    return db.Exec(query).Error
}

func GetLastAppliedMigration(db *gorm.DB) (*Migration, error) {
	var migration Migration
	err := db.Order("applied_at desc").First(&migration).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get last applied migration: %w", err)
	}
	return &migration, nil
}

func MarkMigrationAsRolledBack(db *gorm.DB, migration Migration) error {
	now := time.Now()
	migration.RolledBackAt = &now
	return db.Save(&migration).Error
}

func MarkMigrationAsApplied(db *gorm.DB, migration Migration) error {
	now := time.Now()
	migration.AppliedAt = &now // Ví dụ: thời gian thực thi
	return db.Save(&migration).Error
}
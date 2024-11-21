package migrations

import (
	"gorm.io/gorm"
)

// Up thực thi migration để thêm cột mới vào bảng users
func Up(db *gorm.DB) error {
	// Thêm cột "email_verified" kiểu boolean vào bảng users
	if err := db.Migrator().AddColumn(&User{}, "email_verified"); err != nil {
		return err
	}

	// Hoặc sử dụng để thay đổi cấu trúc bảng tùy chỉnh:
	// if err := db.Exec("ALTER TABLE users ADD COLUMN email_verified BOOLEAN DEFAULT FALSE").Error; err != nil {
	//     return err
	// }

	return nil
}

// Down thực thi rollback migration để xóa cột khỏi bảng users
func Down(db *gorm.DB) error {
	// Xóa cột "email_verified" khỏi bảng users
	if err := db.Migrator().DropColumn(&User{}, "email_verified"); err != nil {
		return err
	}

	// Hoặc sử dụng để thay đổi cấu trúc bảng tùy chỉnh:
	// if err := db.Exec("ALTER TABLE users DROP COLUMN email_verified").Error; err != nil {
	//     return err
	// }

	return nil
}

// User là mô hình của bảng users
type User struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"type:varchar(100)"`
	Email         string `gorm:"type:varchar(100);unique"`
	EmailVerified bool   `gorm:"default:false"` // Cột mới sẽ được thêm vào
}

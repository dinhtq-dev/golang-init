package commands

import (
	"fmt"
	"golang/database/migrations"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// Tạo command "create-migration"
func NewCreateMigrationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-migration [name]",
		Short: "Create a new migration file",
		Args:  cobra.ExactArgs(1), // Chỉ cho phép 1 đối số là tên migration
		Run: func(cmd *cobra.Command, args []string) {
			migrationName := args[0] // Lấy tên migration từ đối số

			// Lấy thời gian hiện tại để tạo tiền tố timestamp
			timestamp := time.Now().Format("20060102150405")

			// Tạo tên file migration theo định dạng timestamp_name.go
			fileName := fmt.Sprintf("database/migrations/%s_%s.go", timestamp, migrationName)

			// Tạo file mới
			file, err := os.Create(fileName)
			if err != nil {
				fmt.Println("Error creating migration file:", err)
				return
			}
			defer file.Close()

			// Nội dung mặc định cho file migration
			migrationContent := fmt.Sprintf(`package migrations

			import "gorm.io/gorm"

			// %s Migration
			func Up(db *gorm.DB) error {
				// Write migration logic here
				return nil
			}

			func Down(db *gorm.DB) error {
				// Write rollback logic here
				return nil
			}
			`, migrationName)

			// Ghi nội dung vào file
			_, err = file.WriteString(migrationContent)
			if err != nil {
				fmt.Println("Error writing to migration file:", err)
				return
			}

			// Thông báo thành công
			fmt.Printf("Migration file created successfully: %s\n", fileName)
		},
	}

	return cmd
}

func NewMigrateDownCommand(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-down",
		Short: "Roll back the last migration",
		Run: func(cmd *cobra.Command, args []string) {
			// Đảm bảo bảng migrations tồn tại
			if err := migrations.EnsureMigrationsTableExists(db); err != nil {
				fmt.Println("Error ensuring migrations table exists:", err)
				return
			}

			// Lấy migration cuối cùng đã chạy
			// lastMigration, err := migrations.GetLastAppliedMigration(db)
			// if err != nil {
			// 	fmt.Println("Error getting last migration:", err)
			// 	return
			// }

			// fmt.Printf("Rolling back migration: %s\n", lastMigration)
			// migrationList := migrations.GetMigrations()
			// for _, migration := range migrationList {
			// 	if migration.Name == lastMigration {
			// 		if err := migration.Down(db); err != nil {
			// 			fmt.Println("Error rolling back migration:", err)
			// 			return
			// 		}
			// 		if err := migrations.arkMigrationAsRolledBack(db, migration.Name); err != nil {
			// 			fmt.Println("Error marking migration as rolled back:", err)
			// 			return
			// 		}
			// 		break
			// 	}
			// }
		},
	}
	return cmd
}



func NewMigrateUpCommand(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-up",
		Short: "Run migrations up",
		Run: func(cmd *cobra.Command, args []string) {
			// Đảm bảo bảng migrations tồn tại
			if err := migrations.EnsureMigrationsTableExists(db); err != nil {
				fmt.Println("Error ensuring migrations table exists:", err)
				return
			}

			// Nạp tất cả các migration từ thư mục
			migrationList, err := migrations.LoadMigrations("database/migrations")
			if err != nil {
				fmt.Println("Error loading migrations:", err)
				return
			}

			log.Println("Migration list:", migrationList)

			// Thực hiện các migration chưa chạy
			for _, migration := range migrationList {
				fmt.Printf("Applying migration: %s\n", migration.Name)
				if err := migration.Up(db); err != nil {
					fmt.Println("Error applying migration:", err)
					return
				}
				// Đánh dấu đã chạy
				// if err := migrations.MarkMigrationAsApplied(db, migration); err != nil {
				// 	fmt.Println("Error marking migration as applied:", err)
				// 	return
				// }
			}
		},
	}
	return cmd
}


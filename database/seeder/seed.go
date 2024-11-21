package seeder

import (
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

// ExecuteCommandSeeder handles the seeding commands
func ExecuteCommandSeeder(db *gorm.DB) {
	if len(os.Args) > 1 {
		// Tách các tham số dòng lệnh
		command := os.Args[1]

		switch command {
		case "--seed":
			var class string
			// Kiểm tra xem có thêm tham số --class hay không
			for i, arg := range os.Args {
				if strings.HasPrefix(arg, "--class=") {
					class = strings.TrimPrefix(arg, "--class=")
					break
				} else if arg == "--class" && i+1 < len(os.Args) {
					class = os.Args[i+1]
					break
				}
			}

			if class != "" {
				runSpecificSeeder(db, class)
			} else {
				runSeeders(db) // Chạy tất cả seeders nếu không chỉ định class
			}

			os.Exit(0) // Thoát sau khi chạy lệnh
		default:
			log.Fatalf("Unknown command: %s", command)
		}
	} else {
		log.Println("No command provided")
	}
}

func runSeeders(db *gorm.DB) {
	log.Println("Running all seeders...")
	SeedUsers(db)
	// Thêm các hàm seed khác nếu có
	log.Println("Seeding completed!")
}

func runSpecificSeeder(db *gorm.DB, class string) {
	log.Printf("Running seeder for class: %s\n", class)
	switch class {
	case "user":
		SeedUsers(db)
	// Thêm các class khác nếu cần
	default:
		log.Printf("Seeder class '%s' not found\n", class)
	}
	log.Println("Seeding completed!")
}

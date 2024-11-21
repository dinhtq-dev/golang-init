// app/main.go
package main

import (
	"fmt"
	"golang/app/container"
	"golang/app/routes"
	"golang/database"
	"golang/database/seeder"
	"golang/langs"

	"golang/config"
)

func main() {
	// Load the configuration
	config, _ := config.Config(".env")

	// Set the current language
	langs.CurrentLanguage = config.AppLang

	// Initialize the database
	db := database.InitDB(config)

	// Initialize the container
	container := container.NewContainer(db)

	// Initialize the router
	r := routes.SetupRouter(container)
	for _, route := range r.Routes() {
		fmt.Printf("%-10s %-35s %s\n", route.Method, route.Path, route.Handler)
	}

	// cmd.RootCmd.AddCommand(commands.NewMigrateUpCommand(db))

	// Run the command
	// cmd.RootCmd.Execute()

	// Run the seeders
	seeder.ExecuteCommandSeeder(db)

	// Run the migrations
	// migrations.ExecuteCommandMigration()

	// r.Run(":" + config.AppPort)
}

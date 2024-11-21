package cmd

import (
	"golang/cmd/commands"

	"github.com/spf13/cobra"
)

// Tạo một root command
var RootCmd = &cobra.Command{
	Use:   "golang",        // Tên ứng dụng
	Short: "MyApp is a CLI tool", // Mô tả ngắn gọn về ứng dụng
}

func init() {
	RootCmd.AddCommand(commands.NewServeCommand()) 
	RootCmd.AddCommand(commands.NewCreateMigrationCommand())

	
}

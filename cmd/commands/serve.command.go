package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Tạo một command "serve"
func NewServeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the application server",
		Run: func(cmd *cobra.Command, args []string) {
			// Logic khi chạy command "serve"
			fmt.Printf("======================================")

			os.Exit(1)
		},
	}

	// Thêm flag cho command "serve"
	// cmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")

	return cmd
}

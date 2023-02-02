package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "list environment variables for aws-assume",
	Run: func(cmd *cobra.Command, args []string) {
		getConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func getConfig() {
	fmt.Printf("AWS_ASSUME_CONFIG_FILE: %s\n", os.Getenv("AWS_ASSUME_CONFIG_FILE"))
	fmt.Printf("AWS_ASSUME_PROFILE: %s\n", os.Getenv("AWS_ASSUME_PROFILE"))
	fmt.Printf("AWS_ASSUME_SECRET_ACCESS_KEY: %s\n", os.Getenv("AWS_ASSUME_SECRET_ACCESS_KEY"))
	fmt.Printf("AWS_ASSUME_ACCESS_KEY_ID: %s\n", os.Getenv("AWS_ASSUME_ACCESS_KEY_ID"))
}

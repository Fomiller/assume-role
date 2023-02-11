package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "print the current environment variables for aws-assume",
	Run: func(cmd *cobra.Command, args []string) {
		getConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func getConfig() {
	fmt.Printf("AWS_ASSUME_CONFIG_DIR: %s\n", AppConfig.GetString("config_dir"))
	fmt.Printf("AWS_ASSUME_PROFILE: %s\n", AppConfig.GetString("profile"))
	fmt.Printf("AWS_ASSUME_SECRET_ACCESS_KEY: %s\n", AppConfig.GetString("secret_access_key"))
	fmt.Printf("AWS_ASSUME_ACCESS_KEY_ID: %s\n", os.Getenv("access_key_id"))
}

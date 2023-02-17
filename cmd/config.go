package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "print the current environment variables for aws-assume",
	Run: func(cmd *cobra.Command, args []string) {
		printConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func printConfig() {
	for _, v := range AppConfig.AllKeys() {
		envName := fmt.Sprintf("AWS_ASSUME_%s", strings.ToUpper(v))
		envValue := AppConfig.GetString(v)
		fmt.Printf("%s: %s\n", envName, envValue)
	}
}

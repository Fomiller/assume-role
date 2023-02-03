package cmd

import (
	"fmt"

	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		printProfileConfig()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&accountFlag, "account", "a", "", "List profiles that have matching account")
	listCmd.Flags().StringVarP(&roleFlag, "role", "r", "", "List profiles that have matching roles")
}

func printProfileConfig() error {
	tFile, err := toml.LoadFile(AppConfig.GetString("config_file"))
	if err != nil {
		return err
	}
	for k, v := range tFile.ToMap() {
		fmt.Printf("[%s]\n", k)
		fmt.Printf("account: %s\n", v.(map[string]interface{})["account"].(string))
		fmt.Printf("role: %s\n\n", v.(map[string]interface{})["role"].(string))
	}
	return nil
}

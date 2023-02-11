package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := printProfileConfig()
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&accountFlag, "account", "a", "", "List profiles that have matching account")
	listCmd.Flags().StringVarP(&roleFlag, "role", "r", "", "List profiles that have matching roles")
}

func printProfileConfig() error {
	configFile := fmt.Sprintf("%s/%s", AppConfig.GetString("config_dir"), DefaultConfigFile)
	profiles, err := ini.Load(configFile)
	if err != nil {
		return err
	}
	for _, v := range profiles.Sections()[1:] {
		fmt.Printf("[%s]\n", v.Name())
		fmt.Printf("account: %s\n", v.Key("account"))
		fmt.Printf("role: %s\n\n", v.Key("role"))
	}
	return nil
}

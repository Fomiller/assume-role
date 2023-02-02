package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := getAssumeConfig()
		if err != nil {
			panic(err)
		}
		printAssumeConfig(cfg)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&accountFlag, "account", "a", "", "List profiles that have matching account")
	listCmd.Flags().StringVarP(&roleFlag, "role", "r", "", "List profiles that have matching roles")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getAssumeConfig() (*ini.File, error) {
	cfg, err := ini.Load("./.aws-assume.toml")
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func printAssumeConfig(cfg *ini.File) {
	for _, section := range cfg.Sections()[1:] {
		fmt.Printf("[%s]\n", section.Name())
		for _, v := range section.Keys() {
			fmt.Printf("%s = %s\n", v.Name(), v.Value())
		}
		fmt.Printf("\n")
	}
	if accountFlag != "" {
		// print accounts that match
	}

	if roleFlag != "" {
		// print accounts that match
	}
}

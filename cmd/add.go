package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a profile to your aws-assume.toml file",
	Run: func(cmd *cobra.Command, args []string) {
		err := addProfile()
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&profileFlag, "profile", "p", "", "Name of new profile you want to configure.")
	addCmd.Flags().StringVarP(&accountFlag, "account", "a", "", "Account Number of the new profile.")
	addCmd.Flags().StringVarP(&roleFlag, "role", "r", "", "Role of the new profile.")
	addCmd.MarkFlagRequired("profile")
	addCmd.MarkFlagRequired("account")
	addCmd.MarkFlagRequired("account")
}

func addProfile() error {
	assumeCfg, err := getAssumeConfig()
	assumeCfg.Section(profileFlag).Key("account").SetValue(accountFlag)
	assumeCfg.Section(profileFlag).Key("role").SetValue(roleFlag)
	err = assumeCfg.SaveTo(assumeConfigPath())
	if err != nil {
		return err
	}

	printSuccessfulCreateProfileMessage()

	return nil
}

func printSuccessfulCreateProfileMessage() {
	fmt.Printf("*******************************************\n\n")
	fmt.Printf("Successfully created Profile: %s\n", profileFlag)
	fmt.Printf("Account: %s\n", accountFlag)
	fmt.Printf("Role: %s\n", roleFlag)
	fmt.Printf("\n*******************************************\n")
}

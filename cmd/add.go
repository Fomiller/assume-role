/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
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
		addProfile()
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
	ProfileConfig.Set(profileFlag, map[string]string{
		"account": accountFlag,
		"role":    roleFlag,
	})
	err := ProfileConfig.WriteConfig()
	if err != nil {
		return err
	}

	fmt.Printf("*******************************************\n")
	fmt.Printf("Successfully created Profile: %s\n", profileFlag)
	fmt.Printf("Account: %s\n", accountFlag)
	fmt.Printf("Role: %s\n", roleFlag)
	fmt.Printf("*******************************************\n")

	return nil
}

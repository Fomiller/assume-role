/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
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
	cfg, err := ini.Load("./.aws-assume.toml")
	if err != nil {
		return err
	}

	newProfileSection, err := cfg.NewSection(profileFlag)
	if err != nil {
		return err
	}
	newAccountKey, err := newProfileSection.NewKey("account", accountFlag)
	if err != nil {
		return err
	}
	newRoleKey, err := newProfileSection.NewKey("role", roleFlag)
	if err != nil {
		return err
	}
	cfg.SaveTo("./.aws-assume.toml")

	fmt.Printf("*******************************************\n")
	fmt.Printf("Successfully created Profile: %s\n", newProfileSection.Name())
	fmt.Printf("Account: %s\n", newAccountKey)
	fmt.Printf("Role: %s\n", newRoleKey)
	fmt.Printf("*******************************************\n")

	return nil
}

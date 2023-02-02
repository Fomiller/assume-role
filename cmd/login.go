/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := login()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func login() error {

	awsRoleArn := os.Getenv("AWS_ROLE_ARN")
	region := "us-east-1"

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return err
	}

	stsClient := sts.NewFromConfig(cfg)
	provider := stscreds.NewAssumeRoleProvider(stsClient, awsRoleArn)
	cfg.Credentials = aws.NewCredentialsCache(provider)
	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("credentials will expire at: ", creds.Expires)

	// var assumeCfg = assumeConfig{}

	credentialsFile, err := ini.Load("~/.aws/credentials")
	if err != nil {
		return err
	}

	credentialsFile.Section(os.Getenv("AWS_ASSUME_PROFILE")).Key("aws_access_key_id").SetValue(creds.AccessKeyID)
	credentialsFile.Section(os.Getenv("AWS_ASSUME_PROFILE")).Key("aws_secret_access_key").SetValue(creds.AccessKeyID)
	credentialsFile.Section(os.Getenv("AWS_ASSUME_PROFILE")).Key("aws_session_token").SetValue(creds.AccessKeyID)
	credentialsFile.SaveTo("$HOME/.aws/credentials")
	return nil
}

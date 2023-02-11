package cmd

import (
	"context"
	"fmt"

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
	Short: "Assume the role defined in the profile that you are passing",
	Long: `Assume role and write credentials including session token for assumed
	role to defined aws credentials file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := login()
		if err != nil {
			cobra.CheckErr(err)
		}
		// viper.SetConfigType("toml")

		// err = login()
		// if err != nil {
		// 	cobra.CheckErr(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&profileFlag, "profile", "p", "", "Name of new profile you want to configure.")
}

func login() error {
	profile := ProfileConfig.AllSettings()[profileFlag].(map[string]interface{})
	arn := createRoleArn(profile["account"].(string), profile["role"].(string))

	awsCreds, err := getCredentials(arn)
	if err != nil {
		return err
	}
	fmt.Println("credentials will expire at: ", awsCreds.Expires)

	credentialsFile := fmt.Sprintf("%s/%s", AppConfig.GetString("credentials_dir"), "credentials")
	creds, err := ini.Load(credentialsFile)
	if err != nil {
		return err
	}

	creds.Section(AppConfig.GetString("profile")).Key("aws_access_key_id").SetValue(awsCreds.AccessKeyID)
	creds.Section(AppConfig.GetString("profile")).Key("aws_secret_access_key").SetValue(awsCreds.SecretAccessKey)
	creds.Section(AppConfig.GetString("profile")).Key("aws_session_token").SetValue(awsCreds.SessionToken)
	creds.SaveTo(credentialsFile)
	return nil
}

func getCredentials(awsRoleArn string) (aws.Credentials, error) {
	region := "us-east-1"

	ctx := context.TODO()

	defaultConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return aws.Credentials{}, err
	}

	stsClient := sts.NewFromConfig(defaultConfig)
	provider := stscreds.NewAssumeRoleProvider(stsClient, awsRoleArn)
	defaultConfig.Credentials = aws.NewCredentialsCache(provider)
	creds, err := defaultConfig.Credentials.Retrieve(context.Background())
	if err != nil {
		return aws.Credentials{}, err
	}

	return creds, nil

}

func createRoleArn(account string, role string) string {
	arn := fmt.Sprintf("arn:aws:iam::%s:role/%s", account, role)
	return arn
}

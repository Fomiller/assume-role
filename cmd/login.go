package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
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
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&profileFlag, "profile", "p", "", "Name of new profile you want to configure.")
}

func login() error {
	assumeCfg, err := getAssumeConfig()
	if err != nil {
		return err
	}

	profile, err := assumeCfg.GetSection(profileFlag)
	if err != nil {
		return err
	}

	arn := createRoleArn(profile.Key("account").String(), profile.Key("role").String())

	awsCreds, err := getCredentials(arn)
	if err != nil {
		return err
	}

	credPath := fmt.Sprintf("%s/%s", AppConfig.GetString("credentials_dir"), DefaultAWSCredentialFile)
	credsDir := path.Dir(credPath)
	err = os.MkdirAll(credsDir, 0755)
	if err != nil {
		return err
	}

	credCfg, err := ini.LooseLoad(credPath)
	if err != nil {
		return err
	}

	updateCredentials(credCfg, credPath, awsCreds, arn)
	printSucessfulAssumeMessage(arn, credPath, awsCreds)

	return nil
}

func getCredentials(arn string) (types.Credentials, error) {
	ctx := context.TODO()
	region := AppConfig.GetString("region")
	accessKeyId := AppConfig.GetString("access_key_id")
	secretAccessKey := AppConfig.GetString("secret_access_key")
	sessionToken := ""

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKeyId,
			secretAccessKey,
			sessionToken,
		)))
	if err != nil {
		return types.Credentials{}, err
	}

	durationSeconds := int32(6 * 60 * 60)
	params := &sts.AssumeRoleInput{
		RoleArn:         &arn,
		RoleSessionName: aws.String("assume-role"),
		DurationSeconds: &durationSeconds,
	}

	stsClient := sts.NewFromConfig(cfg)
	resp, err := stsClient.AssumeRole(context.TODO(), params)
	if err != nil {
		panic("error assuming role: " + err.Error())
		return types.Credentials{}, err
	}
	// provider := stscreds.NewAssumeRoleProvider(stsClient, arn)
	// creds, err := aws.NewCredentialsCache(provider).Retrieve(ctx)
	// if err != nil {
	// 	return aws.Credentials{}, err
	// }

	creds := resp.Credentials
	return *creds, nil

}

func createRoleArn(account string, role string) string {
	arn := fmt.Sprintf("arn:aws:iam::%s:role/%s", account, role)
	return arn
}

func updateCredentials(credFile *ini.File, credPath string, awsCreds types.Credentials, arn string) {
	credFile.Section(AppConfig.GetString("profile")).Key("aws_access_key_id").SetValue(*awsCreds.AccessKeyId)
	credFile.Section(AppConfig.GetString("profile")).Key("aws_secret_access_key").SetValue(*awsCreds.SecretAccessKey)
	credFile.Section(AppConfig.GetString("profile")).Key("aws_session_token").SetValue(*awsCreds.SessionToken)

	credFile.SaveTo(credPath)
}

func printSucessfulAssumeMessage(arn string, credPath string, awsCreds types.Credentials) {
	fmt.Printf("**********************************************\n\n")
	fmt.Printf("Assumed Role %s \n", arn)
	fmt.Printf("Credentials set for [ %s ] profile.\n", AppConfig.GetString("profile"))
	fmt.Printf("Credentials stored in %s.\n", credPath)
	fmt.Printf("Credentials will expire at: %s\n\n", awsCreds.Expiration)
}

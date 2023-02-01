package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	// awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	// awsSecretAccessKeyId := os.Getenv("AWS_SECRET_ACCESS_KEY")
	// awsAccount := os.Getenv("AWS_ACCOUNT")
	// awsDefaultProfile := os.Getenv("AWS_DEFAULT_PROFILE")
	awsRoleArn := os.Getenv("AWS_ROLE_ARN")
	// awsDefaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	// awsSharedCredentialsFile := os.Getenv("AWS_SHARED_CREDENTIALS_FILE")
	region := "us-east-1"

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		panic(err)
	}

	// awsRoleSessionName := "default"
	// stsOptions := sts.Options{Region: "us-east-1"}
	// assumeRoleInput := sts.AssumeRoleInput{
	// 	RoleArn:         &awsRoleArn,
	// 	RoleSessionName: &awsRoleSessionName,
	// 	DurationSeconds: aws.Int32(3600),
	// }

	stsClient := sts.NewFromConfig(cfg)
	provider := stscreds.NewAssumeRoleProvider(stsClient, awsRoleArn)
	// fmt.Println(provider)
	cfg.Credentials = aws.NewCredentialsCache(provider)
	// fmt.Println(cfg)
	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(creds.AccessKeyID)
	fmt.Println(creds.SecretAccessKey)
	fmt.Println(creds.SessionToken)

	os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)
	// res, err := stsClient.AssumeRole(ctx, &assumeRoleInput)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)

	// fmt.Println(awsAccessKeyId)
	// fmt.Println(awsSecretAccessKeyId)
	// fmt.Println(awsAccount)
	// fmt.Println(awsDefaultProfile)
	// fmt.Println(awsRoleArn)
	// fmt.Println(awsDefaultRegion)
	// fmt.Println(awsSharedCredentialsFile)

}

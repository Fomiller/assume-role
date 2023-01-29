package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKeyId := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsAccount := os.Getenv("AWS_ACCOUNT")
	awsDefaultProfile := os.Getenv("AWS_DEFAULT_PROFILE")
	awsRoleArn := os.Getenv("AWS_ROLE_ARN")
	awsDefaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	awsSharedCredentialsFile := os.Getenv("AWS_SHARED_CREDENTIALS_FILE")

	awsRoleSessionName := "default"
	// stsOptions := sts.Options{Region: "us-east-1"}
	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         &awsRoleArn,
		RoleSessionName: &awsRoleSessionName,
		DurationSeconds: aws.Int32(3600),
	}

	ctx := context.TODO()
	sess := session.Must(session.NewSession())
	client := sts.New(sess)

	res, err := client.AssumeRole(ctx, &assumeRoleInput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	fmt.Println(awsAccessKeyId)
	fmt.Println(awsSecretAccessKeyId)
	fmt.Println(awsAccount)
	fmt.Println(awsDefaultProfile)
	fmt.Println(awsRoleArn)
	fmt.Println(awsDefaultRegion)
	fmt.Println(awsSharedCredentialsFile)

}

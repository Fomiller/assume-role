package main

import (
	"github.com/Fomiller/assume-role/cmd"
)

type assumeConfig struct {
	account string
	role    string
}

// func main() {
// awsRoleArn := os.Getenv("AWS_ROLE_ARN")
// region := "us-east-1"
//
// ctx := context.TODO()
// cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
// if err != nil {
// 	panic(err)
// }
//
// stsClient := sts.NewFromConfig(cfg)
// provider := stscreds.NewAssumeRoleProvider(stsClient, awsRoleArn)
// cfg.Credentials = aws.NewCredentialsCache(provider)
// creds, err := cfg.Credentials.Retrieve(context.Background())
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println("credentials will expire at: ", creds.Expires)
//
// os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
// os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)
// os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)
//
// // var assumeCfg = assumeConfig{}
// assumeCfg, err := ini.Load("./.aws-assume.toml")
// if err != nil {
// 	panic(err)
// }
// fmt.Println(assumeCfg.Section("ci-role").Key("account"))
// fmt.Println(assumeCfg.Section("ci-role").Key("role"))

// }

func main() {
	cmd.Execute()
}

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

var (
	profileFlag string
	accountFlag string
	roleFlag    string
	AppConfig   = viper.New()
	HomeDir     string
)

const (
	DefaultAWSCredentialsDir = ".aws"
	DefaultAWSCredentialFile = "credential"
	DefaultConfigFile        = ".aws-assume.ini"
	DefaultProfile           = "default"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "assume-role",
	Short: "switch easily between aws roles",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	ini.DefaultSection = "default"
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.SetEnvPrefix("AWS_ASSUME")
	AppConfig.AutomaticEnv()
	AppConfig.BindEnv("profile")
	AppConfig.BindEnv("region")
	AppConfig.BindEnv("access_key_id")
	AppConfig.BindEnv("secret_access_key")
	AppConfig.SetDefault("credentials_dir", fmt.Sprintf("%s/%s", HomeDir, DefaultAWSCredentialsDir))
	AppConfig.SetDefault("config_dir", HomeDir)
	AppConfig.SetDefault("profile", DefaultProfile)

	defaultRegion, present := os.LookupEnv("AWS_REGION")
	if present {
		AppConfig.SetDefault("region", defaultRegion)
	} else {
		AppConfig.SetDefault("region", "us-east-1")
	}
}

func getAssumeConfig() (*ini.File, error) {
	cfgPath := assumeConfigPath()
	cfg, err := ini.Load(cfgPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func assumeConfigPath() string {
	return fmt.Sprintf("%s/%s", AppConfig.GetString("config_dir"), DefaultConfigFile)
}

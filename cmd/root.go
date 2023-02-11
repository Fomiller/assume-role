package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	profileFlag      string
	accountFlag      string
	roleFlag         string
	AppConfig        = viper.New()
	ProfileConfig    = viper.New()
	CredentialConfig = viper.New()
	HomeDir          string
)

const (
	DefaultAWSCredentialsDir = ".aws/"
	DefaultConfigFile        = ".aws-assume.ini"
	DefaultProfile           = "default"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "assume-role",
	Short: "switch easily between aws roles",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.SetEnvPrefix("AWS_ASSUME")
	AppConfig.AutomaticEnv()
	AppConfig.BindEnv("profile")
	AppConfig.BindEnv("access_key_id")
	AppConfig.BindEnv("secret_access_key")
	AppConfig.SetDefault("credentials_dir", fmt.Sprintf("%s/%s", HomeDir, DefaultAWSCredentialsDir))
	AppConfig.SetDefault("config_dir", HomeDir)
	AppConfig.SetDefault("profile", DefaultProfile)

	ProfileConfig.SetConfigName(".aws-assume")
	ProfileConfig.AddConfigPath(AppConfig.GetString("config_dir"))
	err = ProfileConfig.ReadInConfig()
	if err != nil {
		cobra.CheckErr(err)
	}

	CredentialConfig.SetConfigName("credentials")
	CredentialConfig.SetConfigType("ini")
	CredentialConfig.AddConfigPath(AppConfig.GetString("credentials_dir"))
	err = CredentialConfig.ReadInConfig()
	if err != nil {
		cobra.CheckErr(err)
	}
}

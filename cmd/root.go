package cmd

import (
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
)

const (
	DefaultCredentialsFile = "./credentials"
	DefaultProfile         = "default"
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

	AppConfig.SetEnvPrefix("AWS_ASSUME")
	AppConfig.AutomaticEnv()
	AppConfig.BindEnv("profile")
	AppConfig.BindEnv("credentials_file")
	AppConfig.BindEnv("access_key_id")
	AppConfig.BindEnv("secret_access_key")
	AppConfig.SetDefault("credentials_file", DefaultCredentialsFile)
	AppConfig.SetDefault("profile", DefaultProfile)

	ProfileConfig.SetConfigName(".aws-assume")
	ProfileConfig.AddConfigPath(".")
	err := ProfileConfig.ReadInConfig()
	if err != nil {
		cobra.CheckErr(err)
	}

	CredentialConfig.SetConfigName("credentials")
	CredentialConfig.SetConfigType("ini")
	CredentialConfig.AddConfigPath(".")
	err = CredentialConfig.ReadInConfig()
	if err != nil {
		cobra.CheckErr(err)
	}
}

package internal

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ClientConfig struct {
	BaseUrl  string
	UserName string
	Password string
	LogLevel string
	LogPath  string
}

func ReadConfig(cfg string) {
	if cfg != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfg)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".git-cli" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".git-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetDefault("logPath", "client.log")
	viper.SetDefault("logLevel", "info")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}

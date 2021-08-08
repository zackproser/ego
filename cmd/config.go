package cmd

import (
	"fmt"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getConfigPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Error determining user's home directory for config file")
		return "", err
	}
	path := fmt.Sprintf("/.%s/", strings.ToLower(AppName))
	return fmt.Sprintf("%s/%s", home, path), nil
}

func getConfigName() string {
	return strings.ToLower(AppName)
}

func ReadConfigFile(opts *Options) (*Options, error) {
	fmt.Println("ReadConfigFile")
	viper.SetConfigType("json")
	viper.SetConfigName(getConfigName())

	configPath, err := getConfigPath()
	if err != nil {
		return opts, err
	}
	viper.AddConfigPath(configPath)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Could not find config file for reading")
			return opts, err
		} else {
			fmt.Println("Found and read config")
			fmt.Printf("GithubUsername: %s\n", viper.Get("GithubUsername"))
		}
	}

	loadConfigIntoOptions(opts)

	readGithubTokenFromEnv()

	return opts, nil
}

func readGithubTokenFromEnv() (string, error) {
	if token := viper.GetString("GITHUB_OAUTH_TOKEN"); token != "" {
		return token, nil
	}
	return "", GithubTokenEnvVarUnsetErr{}
}

// loadConfigIntoOptions translates values read from the config file into Options
// consumable by the cmd package
func loadConfigIntoOptions(opts *Options) {
	opts.GithubUsername = viper.GetString("GithubUsername")
}

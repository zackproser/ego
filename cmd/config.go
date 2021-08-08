package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getConfigDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Error determining user's home directory for config file")
		return "", err
	}
	path := fmt.Sprintf("/.%s", strings.ToLower(AppName))
	return fmt.Sprintf("%s/%s", home, path), nil
}

func getConfigName() string {
	return fmt.Sprintf("%s.json", strings.ToLower(AppName))
}

func getConfigPath() (string, error) {
	dir, dirErr := getConfigDir()
	if dirErr != nil {
		logrus.WithFields(logrus.Fields{
			"Error": dirErr,
		}).Debug("Could not determine config dir")
		return "", dirErr
	}

	fileName := getConfigName()

	return filepath.Join(dir, fileName), nil
}

func ReadConfigFile(opts *Options) (*Options, error) {

	viper.SetConfigType("json")
	viper.SetConfigName(getConfigName())

	configPath, err := getConfigDir()
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

	getClient(opts)

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

func configFileExists() bool {
	path, err := getConfigPath()
	fmt.Println("configFileExists: " + path)
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Could not determine path for checking if existing config file present")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("it doesnt")
		return false
	}
	return true
}

func handleConfigCreation() {
	if !configFileExists() {
		fmt.Println("config file does not exist!")
		runSetup()
	}
}

func runSetup() {

	fmt.Println("runSetup")
	var qs = []*survey.Question{
		{
			Name:     "GithubUsername",
			Prompt:   &survey.Input{Message: "Please enter your Github username"},
			Validate: survey.Required,
		},
	}

	answers := struct {
		GithubUsername string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("surveyed: %s\n", answers.GithubUsername)
}

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
	return fmt.Sprintf("%s", strings.ToLower(AppName))
}

func getConfigExt() string {
	return "json"
}

func getConfigPath() (string, error) {
	dir, dirErr := getConfigDir()
	if dirErr != nil {
		logrus.WithFields(logrus.Fields{
			"Error": dirErr,
		}).Debug("Could not determine config dir")
		return "", dirErr
	}

	// Concatenate the filename and extension (.json) to get the second half of the full path
	fileName := fmt.Sprintf("%s.%s", getConfigName(), getConfigExt())

	fmt.Println(filepath.Join(dir, fileName))
	return filepath.Join(dir, fileName), nil
}

func ReadConfigFile(opts *Options) (*Options, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return opts, err
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName(getConfigName())
	viper.SetConfigType("json")

	log.WithFields(logrus.Fields{
		"dir": configDir,
	}).Debug("ReadConfigFile looking for config in directory")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Could not find config file for reading")
			return opts, err
		}
		fmt.Println("Found and read config")
		fmt.Printf("All settings: %+v\n", viper.AllSettings())
	}

	loadConfigIntoOptions(opts)

	readGithubTokenFromEnv()

	instantiateGithubClient(opts)

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
	opts.GithubUsername = viper.GetString("githubusername")
}

func configFileExists() bool {
	path, err := getConfigPath()
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Could not determine path for checking if existing config file present")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Debug("Ego did not find a previously existing config file")
		log.Debug("Ego will now attempt to create one")
		return false
	}
	log.WithFields(logrus.Fields{
		"Path": path,
	}).Debug("Ego found existing config file at path")
	return true
}

func handleConfigCreation() {
	if !configFileExists() {
		err := runSetup()
		if err != nil {
			log.WithFields(logrus.Fields{
				"Error": err,
			}).Debug("Error running config setup")
		} else {
			log.Debug("Successfully setup new configuration")
		}
	}
}

func runSetup() error {

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

	surveyErr := survey.Ask(qs, &answers)
	if surveyErr != nil {
		return surveyErr
	}

	dir, dirErr := getConfigDir()
	if dirErr != nil {
		return dirErr
	}

	viper.AddConfigPath(dir)
	viper.SetConfigName(getConfigName())
	viper.SetConfigType(getConfigExt())

	configPath, configPathErr := getConfigPath()
	if configPathErr != nil {
		return configPathErr
	}

	viper.Set("GithubUsername", answers.GithubUsername)

	_, existErr := os.Stat(configPath)
	if !os.IsExist(existErr) {

		log.WithFields(logrus.Fields{
			"path": configPath,
		}).Debug("Attempting to write config file to path")

		if createDirErr := os.MkdirAll(dir, 0755); createDirErr != nil {
			return createDirErr
		}

		if _, createErr := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0755); createErr != nil {
			return createErr
		}
		log.WithFields(logrus.Fields{
			"path": configPath,
		}).Debug("Created new config file at path")
	}

	configWriteErr := viper.WriteConfigAs(configPath)

	if configWriteErr != nil {
		return configWriteErr
	}
	log.WithFields(logrus.Fields{
		"path": configPath,
	}).Debug("Successfully wrote config file to path")
	return nil
}

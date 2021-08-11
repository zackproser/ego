package cmd

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getSetupQuestions() ([]*survey.Question, *SetupAnswers) {

	questions := []*survey.Question{
		{
			Name:     "GithubUsername",
			Prompt:   &survey.Input{Message: "Please enter your Github username"},
			Validate: survey.Required,
		},
		{
			Name:     "GitRoot",
			Prompt:   &survey.Input{Message: "Enter a directory that contains your git projects"},
			Validate: survey.Required,
		},
	}

	answers := &SetupAnswers{}

	return questions, answers
}

func runSetup() error {

	questions, answers := getSetupQuestions()

	surveyErr := survey.Ask(questions, &answers)
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

	viper.Set("githubusername", answers.GithubUsername)

	// Validate user-supplied path
	if valid, validationErr := pathIsValid(answers.GitRoot); validationErr != nil || !valid {
		return validationErr
	}

	viper.Set("gitroot", answers.GitRoot)

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

package cmd

import "fmt"

type Options struct {
	GithubUsername string
}

func NewOptions() *Options {
	return &Options{}
}

type Configuration struct {
	GithubUsername string
}

type GithubTokenEnvVarUnsetErr struct{}

func (e GithubTokenEnvVarUnsetErr) Error() string {
	return fmt.Sprintf("You must set a valid Github personal access token via $GITHUB_OAUTH_TOKEN")
}

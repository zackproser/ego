package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/v37/github"
)

type Options struct {
	GithubClient   *github.Client
	GithubUsername string
	Tally          *Tally
}

func NewOptions() *Options {
	return &Options{}
}

type Configuration struct {
	GithubUsername string
}
type SetupAnswers struct {
	GithubUsername string
	GitRoot        string
}

type Tally struct {
	repos   []*git.Repository
	commits []*object.Commit
}

// Custom errors
type GithubTokenEnvVarUnsetErr struct{}

func (e GithubTokenEnvVarUnsetErr) Error() string {
	return fmt.Sprintf("You must set a valid Github personal access token via $GITHUB_OAUTH_TOKEN")
}

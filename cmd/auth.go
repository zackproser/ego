package cmd

import (
	"context"

	"github.com/google/go-github/v37/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func getClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.GetString("GITHUB_OAUTH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

func instantiateGithubClient(opts *Options) *Options {
	client := getClient()
	opts.GithubClient = client
	return opts
}

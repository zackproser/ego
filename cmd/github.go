package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/v37/github"
	"github.com/spf13/viper"
)

func getUserPRs(opts *Options) []*github.Issue {
	ctx := context.Background()

	searchOpts := &github.SearchOptions{Sort: "created", Order: "desc"}
	searchString := fmt.Sprintf("is:pr author:%s", viper.Get("githubusername"))
	fmt.Println(searchString)
	sr, _, err := opts.GithubClient.Search.Issues(ctx, searchString, searchOpts)
	if err != nil {
		fmt.Println(err)
	}
	return sr.Issues
}

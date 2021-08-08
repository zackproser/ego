package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/v37/github"
	"github.com/spf13/viper"
)

func getUserPRs(opts *Options) []string {
	ctx := context.Background()

	var pullRequestURLs []string

	searchOpts := &github.SearchOptions{Sort: "created", Order: "desc"}
	searchString := fmt.Sprintf("is:pr author:%s", viper.Get("GithubUsername"))
	sr, _, err := opts.GithubClient.Search.Issues(ctx, searchString, searchOpts)
	if err != nil {
		fmt.Println(err)
	}
	for _, issue := range sr.Issues {
		pullRequestURLs = append(pullRequestURLs, *issue.HTMLURL)
	}
	return pullRequestURLs
}

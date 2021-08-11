package cmd

import (
	"time"

	"github.com/google/go-github/v37/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewTally() *Tally {

	repos, err := getGitReposFromPath(viper.GetString("gitroot"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Error looking up git repos from path")
	}

	return &Tally{
		repos: repos,
	}
}

func (t *Tally) GetRepoCount() int {
	return len(t.repos)
}

func isLastMonth(i *github.Issue) bool {
	month := i.CreatedAt.Month()
	if month == time.July {
		return true
	}
	return false
}

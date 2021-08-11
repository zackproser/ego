package cmd

import (
	"sort"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/v37/github"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewTally instantiates a tally and performs the initial setup to make it ready to use
func NewTally() *Tally {
	repos, err := getGitReposFromPath(viper.GetString("gitroot"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Error looking up git repos from path")
	}

	t := &Tally{
		repos: repos,
	}

	t.BuildCommitMap()

	return t
}

func (t *Tally) GetRepoCount() int {
	return len(t.repos)
}

func (t *Tally) BuildCommitMap() {
	for _, r := range t.repos {
		commitIter, commitErr := r.CommitObjects()
		if commitErr != nil {
			log.Debug("Error getting commit iter")
			continue
		}
		commitErr = commitIter.ForEach(func(c *object.Commit) error {
			t.commits = append(t.commits, c)
			return nil
		})
	}
}

func (t *Tally) GetAuthorMap(desiredLength int) AuthorList {
	authorMap := make(map[string]int)
	for _, c := range t.commits {
		authorMap[c.Author.Name]++
	}
	return rankByCommitCount(authorMap)[:desiredLength]
}

func (t *Tally) FilterCommits(commits []*object.Commit, f func(*object.Commit) bool) []*object.Commit {
	fc := make([]*object.Commit, 0)
	for _, commit := range commits {
		if f(commit) {
			fc = append(fc, commit)
		}
	}
	return fc
}

func (t *Tally) FilterCommitsByAuthorName(name string) []*object.Commit {
	return t.FilterCommits(t.commits, func(c *object.Commit) bool {
		return c.Author.Name == name
	})
}

func (t *Tally) GetAuthorCounts(timePeriod string) (map[string]int, error) {
	// Start by just operating on one repo
	r := t.repos[0]
	commitIter, commitErr := r.CommitObjects()
	if commitErr != nil {
		log.Debug("Error getting commit iter")
		return nil, commitErr
	}

	authorMap := make(map[string]int)

	commitErr = commitIter.ForEach(func(c *object.Commit) error {
		author := c.Author
		authorMap[author.Name]++

		return nil
	})
	return authorMap, nil
}

func isLastMonth(i *github.Issue) bool {
	month := i.CreatedAt.Month()
	if month == time.July {
		return true
	}
	return false
}

type Author struct {
	Name    string
	Commits int
}

type AuthorList []Author

func rankByCommitCount(commitCounts map[string]int) AuthorList {
	al := make(AuthorList, len(commitCounts))
	i := 0
	for k, v := range commitCounts {
		al[i] = Author{k, v}
		i++
	}
	sort.Sort(sort.Reverse(al))
	return al
}

func (a AuthorList) Len() int           { return len(a) }
func (a AuthorList) Less(i, j int) bool { return a[i].Commits < a[j].Commits }
func (a AuthorList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

func normalizePath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		homedir, homedirErr := homedir.Dir()
		if homedirErr != nil {
			log.WithFields(logrus.Fields{
				"Error": homedirErr,
			}).Debug("Error determining user homedir")
			return "", homedirErr
		}
		return filepath.Join(homedir, strings.TrimLeft(path, "~/")), nil
	}
	return path, nil
}

func pathIsValid(path string) (bool, error) {
	normalizedPath, normalizeErr := normalizePath(path)
	if normalizeErr != nil {
		return false, normalizeErr
	}
	if _, err := os.Stat(normalizedPath); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func getGitReposFromPath(path string) ([]*git.Repository, error) {
	normalizedPath, normalizedErr := normalizePath(path)
	if normalizedErr != nil {
		return nil, normalizedErr
	}
	dirEntries, readErr := os.ReadDir(normalizedPath)
	if readErr != nil {
		return nil, readErr
	}
	return filterGitRepos(normalizedPath, dirEntries), nil
}

func filterGitRepos(path string, entries []os.DirEntry) []*git.Repository {
	var validGitRepos []*git.Repository

	for _, dirEntry := range entries {
		fullPath := filepath.Join(path, dirEntry.Name())
		r, err := git.PlainOpen(fullPath)
		if err != nil {
			log.WithFields(logrus.Fields{
				"Error": err,
			}).Error("Could not open git repo")
			continue
		}
		validGitRepos = append(validGitRepos, r)
	}

	return validGitRepos
}

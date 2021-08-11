package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

func pathIsValid(path string) (bool, error) {
	if strings.HasPrefix(path, "~/") {
		homedir, homedirErr := homedir.Dir()
		if homedirErr != nil {
			log.WithFields(logrus.Fields{
				"Error": homedirErr,
			}).Debug("Error determining user homedir")
			return false, homedirErr
		}
		path = filepath.Join(homedir, path[2:])
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

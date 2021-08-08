package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var AppName = "Tracker"

var log = logrus.New()

var rootCmd = &cobra.Command{
	Use:              "tracker",
	Short:            "tracker",
	Long:             "tracker",
	PersistentPreRun: persistentPreRun,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func initConfig() error {
	log.SetLevel(logrus.DebugLevel)

	// Look for a config file in the default storage directory
	opts := NewOptions()
	opts, err := ReadConfigFile(opts)
	if err != nil {
		return err
	}
	fmt.Printf("PROCESSED OPTS: %+v\n", opts)
	return nil
}

func persistentPreRun(cmd *cobra.Command, args []string) {
	fmt.Println("persistentPreRun")
	err := initConfig()
	if err != nil {
		log.Error(err)
	}
}

func Execute() {
	fmt.Println("Execute")
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
		return
	}
}

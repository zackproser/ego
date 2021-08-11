package cmd

import (
	"github.com/pterm/pterm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var AppName = "Ego"

var log = logrus.New()

var rootCmd = &cobra.Command{
	Use:   "ego",
	Short: "ego",
	Long:  "ego",
	Run: func(cmd *cobra.Command, args []string) {
		//renderMarquee()

		handleConfigCreation()

		opts, err := initConfig()
		if err != nil {
			log.Error(err)
		}

		opts = instantiateGithubClient(opts)

		//startSpinner()
		//renderUserPRs(opts)
		dummyStatsOutput(opts)
		//stopSpinner()
	},
}

func init() {
	log.SetLevel(logrus.DebugLevel)
}

func renderUserPRs(opts *Options) {
	prs := getUserPRs(opts)

	renderUI(prs, opts)
}

func stopSpinner() {
	pterm.DefaultSpinner.Stop()
}

func startSpinner() {
	pterm.DefaultSpinner.Start()
}

func initConfig() (*Options, error) {
	log.SetLevel(logrus.DebugLevel)

	// Look for a config file in the default storage directory
	opts := NewOptions()
	opts, err := ReadConfigFile(opts)
	if err != nil {
		return opts, err
	}
	opts.Tally = NewTally()
	return opts, nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
		return
	}
}

package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var AppName = "Tracker"

var log = logrus.New()

var rootCmd = &cobra.Command{
	Use:   "tracker",
	Short: "tracker",
	Long:  "tracker",
	Run: func(cmd *cobra.Command, args []string) {

		handleConfigCreation()

		opts, err := initConfig()
		if err != nil {
			log.Error(err)
		}

		startSpinner()
		renderUserPRs(opts)
		stopSpinner()

	},
}

func renderUserPRs(opts *Options) {
	prs := getUserPRs(opts)
	pterm.DefaultHeader.Println("Pull Requests")
	var items []pterm.BulletListItem
	for _, prUrl := range prs {
		item := pterm.BulletListItem{
			Level: 0,
			Text:  prUrl,
		}
		items = append(items, item)
	}
	pterm.DefaultBulletList.WithItems(items).Render()
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
	return opts, nil
}

func persistentPreRun(cmd *cobra.Command, args []string) {
	fmt.Println("persistentPreRun")
}

func Execute() {
	fmt.Println("Execute")
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
		return
	}
}

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

		handleConfigCreation()

		/*
			_, err := initConfig()
			if err != nil {
				log.Error(err)
			} */

		/*
			startSpinner()
			renderUserPRs(opts)
			stopSpinner()
		*/

	},
}

func init() {
	log.SetLevel(logrus.DebugLevel)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
		return
	}
}

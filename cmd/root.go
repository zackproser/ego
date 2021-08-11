package cmd

import (
	"time"

	"github.com/google/go-github/v37/github"
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
		renderUserPRs(opts)
		//stopSpinner()
	},
}

func renderMarquee() {
	egoLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("E", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(egoLogo)

	pterm.DefaultCenter.Print(pterm.DefaultBasicText.Sprint("Work stats tracker"))

	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightGreen)).WithMargin(10).Sprint("By Zack Proser"))

	time.Sleep(3 * time.Second)

	// Clear the screen
	print("\033[H\033[2J")

}

func init() {
	log.SetLevel(logrus.DebugLevel)
}

func isLastMonth(i *github.Issue) bool {
	month := i.CreatedAt.Month()
	if month == time.July {
		return true
	}
	return false
}

func renderUserPRs(opts *Options) {
	prs := getUserPRs(opts)

	renderUI(prs)
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

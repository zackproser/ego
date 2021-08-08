package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var AppName = "Tracker"

var log = logrus.New()

var rootCmd = &cobra.Command{
	Use:   "tracker",
	Short: "tracker",
	Long:  "tracker",
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	fmt.Println("initConfig")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
		return
	}
}

func main() {
	fmt.Println("vim-go")
}

package cmd

import (
	"time"

	"github.com/pterm/pterm"
)

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

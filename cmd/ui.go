package cmd

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/google/go-github/v37/github"
)

func renderUI(issues []*github.Issue) {
	if initErr := ui.Init(); initErr != nil {
		log.Fatalf("failed to initialize termui: %v", initErr)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to switch tabs"
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue
	header.TextStyle.Fg = ui.ColorWhite

	p2 := widgets.NewParagraph()
	p2.Text = "Press q to quit\nPress h or l to switch tabs\n"
	p2.Title = "Keys"
	p2.SetRect(5, 5, 40, 15)
	p2.BorderStyle.Fg = ui.ColorYellow

	bc := widgets.NewBarChart()
	bc.Title = "Wakka Chart"
	bc.Data = []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(5, 5, 50, 10)
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}

	table1 := widgets.NewTable()
	table1.Rows = [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"2016", "10", "11"},
	}
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(5, 5, 50, 10)

	tabpane := widgets.NewTabPane("Pull Requests", "Issues", "Commits")
	tabpane.SetRect(0, 1, 50, 4)
	tabpane.Border = true

	/*	renderTab := func() {
			switch tabpane.ActiveTabIndex {
			case 0:
				ui.Render(p2)
			case 1:
				ui.Render(bc)
			case 2:
				ui.Render(header, tabpane, table1)
			}
		}

		ui.Render(header, tabpane, p2) */

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, p2),
			ui.NewCol(1.0/2, bc),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, table1),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
			/*case "h":
				tabpane.FocusLeft()
				ui.Clear()
				ui.Render(header, tabpane)
				renderTab()
			case "l":
				tabpane.FocusRight()
				ui.Clear()
				ui.Render(header, tabpane)
				renderTab()*/
		}
	}
}

func filterLastMonthsPRs(issues []*github.Issue) {
	var lastMonthPRs []*github.Issue

	for _, pr := range issues {
		if isLastMonth(pr) {
			lastMonthPRs = append(lastMonthPRs, pr)
		}
	}
}

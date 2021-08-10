package cmd

import (
	"fmt"
	"strings"

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
	header.SetRect(0, 0, 5, 10)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue
	header.TextStyle.Fg = ui.ColorWhite

	bc := widgets.NewBarChart()
	bc.Title = "Wakka Chart"
	bc.Data = []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(5, 5, 50, 10)
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}

	table1 := widgets.NewTable()
	table1.Rows = [][]string{
		[]string{"Repo", "PR Title", "URL"},
		[]string{issues[0].GetRepository().GetName(), issues[0].GetTitle(), issues[0].GetHTMLURL()},
		[]string{issues[1].GetRepository().GetName(), issues[1].GetTitle(), issues[1].GetHTMLURL()},
		[]string{issues[2].GetRepository().GetName(), issues[2].GetTitle(), issues[2].GetHTMLURL()},
		[]string{issues[3].GetRepository().GetName(), issues[3].GetTitle(), issues[3].GetHTMLURL()},
	}
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 0, 60, 10)
	table1.RowSeparator = true
	table1.FillRow = true

	termWidth, termHeight := ui.TerminalDimensions()

	tabpane := widgets.NewTabPane("Pull Requests", "Issues", "Commits")
	tabpane.SetRect(0, 0, 50, 1)
	tabpane.Border = false

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(table1)
		case 1:
			ui.Render(table1)
		}
	}

	grid := ui.NewGrid()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(.9/3, tabpane),
			ui.NewCol(1.2/3, table1),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch key := e.ID; key {
		case "<Esc>", "<C-c>":
			return
		case "<Left>":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "<Right>":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
			fmt.Println("Key is: " + key)
			filteredIssues := filterIssuesByTypeaheadSearch(issues, key)
			rowsMap := [][]string{}
			for i, _ := range filteredIssues {
				m := []string{}
				m = append(m, filteredIssues[i].GetRepository().GetName(), filteredIssues[i].GetTitle(), filteredIssues[i].GetHTMLURL())
				rowsMap = append(rowsMap, m)
			}
			if len(rowsMap) == 0 {
				rowsMap = [][]string{
					[]string{"No", "Results", "Found"},
				}
			}
			table1.Rows = rowsMap
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		}

	}
}

func filterIssuesByTypeaheadSearch(issues []*github.Issue, key string) []*github.Issue {
	var filteredIssues []*github.Issue
	for _, issue := range issues {
		if strings.HasPrefix(strings.ToLower(issue.GetTitle()), strings.ToLower(key)) {
			filteredIssues = append(filteredIssues, issue)
		}
	}
	return filteredIssues
}

func filterLastMonthsPRs(issues []*github.Issue) {
	var lastMonthPRs []*github.Issue

	for _, pr := range issues {
		if isLastMonth(pr) {
			lastMonthPRs = append(lastMonthPRs, pr)
		}
	}
}

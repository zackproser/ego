package cmd

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/google/go-github/v37/github"
	"github.com/rivo/tview"
)

func dummyStatsOutput(opts *Options) {
	fmt.Printf("Total Local Repos: %+v\n", opts.Tally.GetRepoCount())

	authorMap := opts.Tally.GetAuthorMap(10)
	fmt.Println("Local Repo Author Map:")
	for k, v := range authorMap {
		fmt.Printf("%s: %v\n", k, v)
	}

	commitsByZack := opts.Tally.FilterCommitsByAuthorName("Zack Proser")
	for _, commit := range commitsByZack {
		fmt.Printf("I wrote: %s", commit.Message)
	}
}

func renderUI(issues []*github.Issue, opts *Options) {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	cols := 2
	i := 0
	for r := 0; r < len(issues); r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			table.SetCell(r, c,
				tview.NewTableCell(issues[i].GetTitle()).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
			table.SetCell(r, c+1,
				tview.NewTableCell(issues[i].GetHTMLURL()).
					SetTextColor(tcell.ColorGreen).
					SetAlign(tview.AlignCenter))
			table.SetCell(r, c+2,
				tview.NewTableCell(issues[i].GetState()).
					SetTextColor(tcell.ColorGreen).
					SetAlign(tview.AlignCenter))

			i = (i + 1) % len(issues)
		}
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
		fmt.Printf("Selected: %+v\n", table.GetCell(row, column).Text)
	})
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	/*table1 := NewTable()
	table1.Rows = [][]string{
		[]string{"Repo", "PR Title", "URL"},
		[]string{issues[0].GetRepository().GetFullName(), issues[0].GetTitle(), issues[0].GetHTMLURL()},
		[]string{issues[1].GetRepository().GetName(), issues[1].GetTitle(), issues[1].GetHTMLURL()},
		[]string{issues[2].GetRepository().GetName(), issues[2].GetTitle(), issues[2].GetHTMLURL()},
		[]string{issues[3].GetRepository().GetName(), issues[3].GetTitle(), issues[3].GetHTMLURL()},
	}

	termWidth, termHeight := ui.TerminalDimensions()

	tabpane := widgets.NewTabPane("Pull Requests", "Issues", "Commits")
	tabpane.SetRect(0, 0, termWidth, termHeight)
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
		ui.NewRow(1.0/6,
			ui.NewCol(1.0/2, header)),
		ui.NewRow(1.0/1.0,
			ui.NewCol(1.0/1.0, table1)),
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
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "<Down>":
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
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

	}*/
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

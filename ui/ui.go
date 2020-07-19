package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type UI struct {
	App     *tview.Application
	Side    *tview.List
	Main    *tview.TextView
	Console *tview.TextView
}

var (
	toggle = true
)

func NewUI() *UI {
	app := tview.NewApplication()
	side := tview.NewList()
	main := tview.NewTextView().SetWrap(true)
	console := tview.NewTextView()

	grid := tview.NewGrid().
		SetRows(0, 1).
		SetColumns(50, 0).
		SetBorders(true)

	grid.AddItem(side, 0, 0, 1, 1, 0, 0, false).
		AddItem(main, 0, 1, 1, 1, 0, 0, false).
		AddItem(console, 1, 0, 1, 2, 0, 0, false)
	app.SetRoot(grid, true)

	changeView := func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if toggle {
				app.SetFocus(main)
				toggle = false
			} else {
				app.SetFocus(side)
				toggle = true
			}
			return nil
		}
		return event
	}
	side.SetInputCapture(changeView)
	main.SetInputCapture(changeView)

	return &UI{
		App:     app,
		Side:    side,
		Main:    main,
		Console: console,
	}
}

func (ui *UI) Run() error {
	return ui.App.Run()
}

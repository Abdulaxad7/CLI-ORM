package src

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

func DbSelect(app *tview.Application) {

	list := tview.NewList()

	list.AddItem("Mysql", "", '1', func() { Msql(app) })
	list.AddItem("Postgresql", "", '2', func() { Psql(app) })
	list.AddItem("Quit", "", '0', func() { exit(app) })

	list.SetBorder(true).SetBorderColor(tcell.ColorGreen)
	if err := app.SetRoot(list, true).SetFocus(list).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func exit(app *tview.Application) {
	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Cancel", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
				os.Exit(0)
			} else {
				DbSelect(app)
			}
		})
	if err := app.SetRoot(modal, true).SetFocus(modal).Run(); err != nil {
		panic(err)
	}
}

package src

import (
	"github.com/rivo/tview"
)

func DbSelect(app *tview.Application) {

	list := tview.NewList()

	list.AddItem("Mysql", "", '1', func() { Msql(app) })
	list.AddItem("Postgres", "", '2', func() { Psql(app) })
	list.AddItem("Quit", "", '0', func() { exit(app) })

	list.SetBorder(true)
	if err := app.SetRoot(list, true).SetFocus(list).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func exit(app *tview.Application) {
	app.Stop()
}

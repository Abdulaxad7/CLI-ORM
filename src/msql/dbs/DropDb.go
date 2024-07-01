package dbs

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func (c CRUD) DeleteForm(app *tview.Application, db *gorm.DB) *tview.Form {
	show := ShowS{}
	deleteDb := tview.NewForm()
	deleteDb.AddInputField("Database name to drop", "", 20, nil, nil)
	deleteDb.AddButton("Delete", func() {
		if dbName := c.CaptureDbName(deleteDb, db, 1); dbName != "" {
			c.Layout(app, db, dbName, "drop", "", 1)
		} else {
			show.ShowDbs(app, db)
		}
	})
	deleteDb.SetBorder(true).SetBorderColor(tcell.ColorGreen).SetTitle("Drop db")
	return deleteDb
}

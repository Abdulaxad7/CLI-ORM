package dbs

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func (c CRUD) CreateDb(app *tview.Application, db *gorm.DB) *tview.Form {
	show := ShowS{}
	form := tview.NewForm()
	form.AddInputField("Database name to create ", "", 20, nil, nil)
	form.AddButton("Create", func() {
		if dbName := c.CaptureDbName(form, db, 0); dbName != "" {
			c.Layout(app, db, dbName, "create", "", 0)
		} else {
			show.ShowDbs(app, db)
		}
	})
	form.SetTitle("Create db").SetBorder(true).SetBorderColor(tcell.ColorGreen)
	return form
}

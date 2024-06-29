package msql

import (
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func CreateDb(app *tview.Application, db *gorm.DB) {
	form := tview.NewForm()

	if err := app.SetRoot(form, true).SetFocus(form); err != nil {
		panic(err)
	}
}

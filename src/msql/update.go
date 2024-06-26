package msql

import (
	"fmt"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func updateValue(app *tview.Application, db *gorm.DB, dbName string, table string, value string) {
	form := tview.NewForm()

	form.AddInputField(fmt.Sprintf("Updating from %s.%s value=%s to", dbName, table, value), "", 30, nil, nil)
	form.AddButton("Update", func() {
		update(db, table)
	})
	form.SetBorder(true)
	app.SetRoot(form, true)
}
func update(db *gorm.DB, table string) {
	db.Exec("SET SQL_SAFE_UPDATES = 0")
	db.Exec(fmt.Sprintf("update %s set %s where price=56.99 and artist='John Coltrane';", table))
}

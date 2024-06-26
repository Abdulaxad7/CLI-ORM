package msql

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func ShowTables(app *tview.Application, db *gorm.DB) {
	tables := tview.NewTable()
	tables.SetBorders(true).SetBorder(true)
	var results []map[string]interface{}
	db.Raw("SHOW TABLES;").Scan(&results)
	for i, v := range results {
		for _, v2 := range v {

			tables.SetCell(i, 0, tview.NewTableCell(v2.(string)).SetTextColor(tcell.ColorGreenYellow).SetAlign(tview.AlignCenter)).
				SetSelectable(true, false).
				SetOffset(1, 1).SetBackgroundColor(tcell.ColorWhite)

		}

	}

	tables.Select(0, 0).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			ShowDbs(app, db)
		}
		if key == tcell.KeyEnter {
			tables.SetSelectable(true, true)
		}
	}).
		SetSelectedFunc(func(row int, column int) {
			v2 := tables.GetCell(row, column).Text
			ShowValues(app, db, v2)
		})

	app.SetRoot(tables, true).SetFocus(tables)
}

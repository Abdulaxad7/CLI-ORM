package msql

import (
	"Cli-Orm/config/mq"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func ShowTables(app *tview.Application, db *gorm.DB, dbName string) {
	tables := tview.NewTable()
	tables.SetBorders(true).SetBorder(true)

	results := mq.DbTables(db)
	for i, v := range results {
		for _, v2 := range v {

			tables.SetCell(i, 0, tview.NewTableCell(v2.(string)).SetTextColor(tcell.ColorGreenYellow).SetAlign(tview.AlignCenter)).
				SetSelectable(true, false).
				SetOffset(1, 1)

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
			valFrom := tables.GetCell(row, column).Text
			ShowValues(app, db, valFrom, dbName)
		})

	if err := app.SetRoot(tables, true).SetFocus(tables).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

package msql

import (
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func ShowValues(app *tview.Application, db *gorm.DB, table string, dbName string) {
	values := tview.NewTable()
	values.SetBorders(true).SetBorder(true).SetTitle(table)
	results := mq.DbValues(db, table)
	var i, j int
	for _, v := range results {
		j = 0
		for _, v2 := range v {

			cellValue := fmt.Sprintf("%v", v2)
			values.SetCell(i, j, tview.NewTableCell(cellValue).SetTextColor(tcell.ColorRed).
				SetAlign(tview.AlignCenter)).
				SetSelectable(true, true)
			j++
		}
		i++
	}

	values.SetDoneFunc(func(key tcell.Key) {

		if key == tcell.KeyEscape {
			ShowTables(app, db, dbName)
		}
		if key == tcell.KeyEnter {
			values.SetSelectable(true, true)
		}

	}).SetSelectedFunc(func(row int, column int) {
		value := values.GetCell(row, column).Text
		updateValue(app, db, dbName, table, value)
	})

	if err := app.SetRoot(values, true).SetFocus(values).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

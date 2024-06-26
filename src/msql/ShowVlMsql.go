package msql

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func ShowValues(app *tview.Application, db *gorm.DB, table string, dbName string) {
	var results []map[string]interface{}
	var myColumns []string

	datas := tview.NewTable()
	datas.SetBorders(true).SetBorder(true)
	db.Raw(fmt.Sprintf("SELECT*FROM %s;", table)).Find(&results)
	db.Raw(fmt.Sprintf("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS  WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s';", dbName, table)).Find(&myColumns)

	var i, j int

	for _, v := range results {
		j = 0
		for _, v2 := range v {
			cellValue := fmt.Sprintf("%v", v2)
			datas.SetCell(i, j, tview.NewTableCell(cellValue).SetTextColor(tcell.ColorRed).SetAlign(tview.AlignCenter)).SetSelectable(true, true)
			j++

		}
		i++
	}

	datas.SetDoneFunc(func(key tcell.Key) {

		if key == tcell.KeyEscape {
			ShowTables(app, db, dbName)
		}
		if key == tcell.KeyEnter {
			datas.SetSelectable(true, true)
		}

	}).SetSelectedFunc(func(row int, column int) {
		value := datas.GetCell(row, column).Text
		updateValue(app, db, dbName, table, value)
	})

	app.SetRoot(datas, true).SetFocus(datas)
}

func columns() {

}

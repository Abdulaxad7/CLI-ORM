package msql

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func ShowValues(app *tview.Application, db *gorm.DB, valFrom string) {
	datas := tview.NewTable()
	datas.SetBorders(true).SetBorder(true)
	var results []map[string]interface{}
	db.Raw(fmt.Sprintf("SELECT*FROM %s;", valFrom)).Find(&results)
	var i, j int

	for _, v := range results {
		j = 0
		for _, v2 := range v {
			cellValue := fmt.Sprintf("%v", v2)
			datas.SetCell(i, j, tview.NewTableCell(cellValue).SetTextColor(tcell.ColorRed).SetAlign(tview.AlignCenter))
			j++

		}
		i++
	}
	datas.SetDoneFunc(func(key tcell.Key) {

		if key == tcell.KeyEscape {
			ShowTables(app, db)
		}

	})
	app.SetRoot(datas, true).SetFocus(datas)
}

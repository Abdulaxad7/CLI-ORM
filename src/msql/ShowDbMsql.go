package msql

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func ShowDbs(app *tview.Application, db *gorm.DB) {

	tb := tview.NewTable()
	tb.SetBorders(true).SetBorder(true)
	var results []map[string]interface{}
	db.Raw("SHOW DATABASES;").Scan(&results)
	for i, v := range results {
		for _, v2 := range v {
			tb.SetCell(i, 0, &tview.TableCell{Text: v2.(string), Align: tview.AlignCenter, Color: tview.Styles.TitleColor}).
				SetSelectable(true, false)
		}
	}
	tb.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {

			tb.SetSelectable(true, true)
		}
	}).
		SetSelectedFunc(func(row int, column int) {
			dbName := tb.GetCell(row, column).SetTextColor(tcell.ColorRed).Text
			db.Exec(fmt.Sprintf("USE %s", dbName))

			ShowTables(app, db, dbName)
		})

	app.SetRoot(tb, true).SetFocus(tb)

}

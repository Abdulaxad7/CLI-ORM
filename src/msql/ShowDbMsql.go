package msql

import (
	"Cli-Orm/config/mq"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
	"os"
)

func (s ShowS) ShowDbs(app *tview.Application, db *gorm.DB) {
	cr := CRUD{}
	var dbName string
	tb := tview.NewTable()

	tb.SetBorders(true).SetBorder(true).SetTitle("Databases").SetBorderColor(tcell.ColorGreen)

	results := mq.DbQuery(db)
	for i, v := range results {
		for _, v2 := range v {
			tb.SetCell(i, 0, &tview.TableCell{Text: v2.(string), Align: tview.AlignCenter, Color: tview.Styles.TitleColor}).
				SetSelectable(true, false)
		}
	}
	tb.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {

		if key == tcell.KeyEscape {
			app.Stop()
			os.Exit(0)
		}
		if key == tcell.KeyEnter {
			tb.SetSelectable(true, true)
		}

	}).SetSelectedFunc(func(row int, column int) {
		dbName = tb.GetCell(row, column).SetTextColor(tcell.ColorRed).Text
		mq.DbUseQuery(db, dbName)
		s.ShowTables(app, db, dbName)
	})

	flex := tview.NewFlex()
	flex.AddItem(tb, 0, 10, true)
	flex.AddItem(cr.DeleteForm(app, db), 0, 5, false)
	flex.AddItem(cr.CreateDb(app, db), 0, 5, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

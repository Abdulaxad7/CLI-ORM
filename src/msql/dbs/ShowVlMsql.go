package dbs

import (
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func (s ShowS) ShowValues(app *tview.Application, db *gorm.DB, table string, dbName string) {
	cr := CRUD{}
	values := tview.NewTable()
	values.SetBorders(true).SetBorder(true).SetTitle(dbName + "/" + table + "/")
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
			s.ShowTables(app, db, dbName)
		}
		if key == tcell.KeyEnter {
			values.SetSelectable(true, true)
		}

	}).SetSelectedFunc(func(row int, column int) {
		value := values.GetCell(row, column).Text
		cr.UpdateValue(app, db, dbName, table, value)
	})
	values.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'i' {
			InsertIntoTb(app, db, dbName, table)
		}
		if event.Rune() == 'd' {
			cr.Layout(app, db, dbName, table, "drop", 2)
		}
		return event
	})
	flex := tview.NewFlex()
	flex.AddItem(values, 0, 12, true)
	flex.AddItem(Info("\n\n Press ⮐ enter to change value\n\n"+
		" Press ⎋esc to go back\n\n"+
		" ↑/ up • ↓/ down\n\n"+
		" ←/ right • →/ left\n\n"+
		" Press 'i' to insert values\n\n"+
		" Press 'd' to drop table "),
		0, 2, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

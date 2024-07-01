package dbs

import (
	"Cli-Orm/config/mq"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func (s ShowS) ShowTables(app *tview.Application, db *gorm.DB, dbName string) {
	tables := tview.NewTable()
	tables.SetBorders(true).SetBorder(true).SetTitle(dbName + "/")

	results := mq.DbTables(db)
	for i, v := range results {
		for _, v2 := range v {

			tables.SetCell(i, 0, tview.NewTableCell(v2.(string)).SetTextColor(tcell.ColorGreenYellow).SetAlign(tview.AlignCenter)).
				SetSelectable(true,
					false).
				SetOffset(1, 1)
		}
	}
	tables.Select(0, 0).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			s.ShowDbs(app, db)
		}
		if key == tcell.KeyEnter {
			tables.SetSelectable(true, true)
		}
	}).
		SetSelectedFunc(func(row int, column int) {
			valFrom := tables.GetCell(row, column).Text
			s.ShowValues(app, db, valFrom, dbName)
		})
	tables.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'c' {
			CreateTb(app, db, dbName)
		}
		return event
	})
	flex := tview.NewFlex()
	flex.AddItem(tables, 0, 12, true)
	flex.AddItem(Info("\n\n Press ⮐ enter to show raws\n\n"+
		" Press ⎋esc to go back\n\n"+
		" ↑/ up • ↓/ down\n\n"+
		" Press 'c' to create table\n\n"),
		0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

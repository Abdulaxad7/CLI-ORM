package dbs

import (
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func InsertIntoTb(app *tview.Application, db *gorm.DB, dbName, tbName string) {
	form := tview.NewForm()
	columns := CutMapStringInterface(mq.DbTableColumns(db, dbName, tbName))
	columnTypes := CutMapStringInterface(mq.DbDataTypes(db, dbName, tbName))
	for i, column := range columns {
		form.AddInputField(fmt.Sprintf("%s (%s)", column, columnTypes[i]), "", 20, nil, nil)
	}
	form.AddButton("Insert", func() {
		values := checkInserted(form, len(columns))
		mq.DbInsertToTable(db, tbName, values)
	})
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Rune() == 'g' {
			ShowS{}.ShowTables(app, db, dbName)
		}
		return event
	})
	form.SetBorder(true).SetTitle(dbName + "/" + tbName + "/").SetBorderColor(tcell.ColorGreen)
	flex := tview.NewFlex()
	flex.AddItem(form, 0, 12, true)
	flex.AddItem(Info("\n\n Press ⎋esc or 'g' to go back\n\n"), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
func checkInserted(form *tview.Form, size int) []string {
	var result []string
	for i := 0; i < size; i++ {
		result = append(result, form.GetFormItem(i).(*tview.InputField).GetText())
	}
	return result
}

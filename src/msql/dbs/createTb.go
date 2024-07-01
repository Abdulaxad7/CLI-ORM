package dbs

import (
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
	"strconv"
	"unicode"
)

var s = ShowS{}

func (s ShowS) CreateTb(app *tview.Application, db *gorm.DB, dbName string) {
	form := tview.NewForm()
	var columns string
	var tableName string
	form.AddInputField("Column number ", "", 20, nil, func(text string) {
		columns = text
	})
	form.AddInputField("Table Name", "", 20, nil, func(text string) {
		tableName = text
	})

	form.AddButton("Create", func() {
		tables := mq.DbTables(db)
		if tableName == "" || columns == "" || ContainsAlpha(columns) || ContainsString(tables, tableName) || unicode.IsDigit(rune(tableName[0])) {
			s.CreateTb(app, db, dbName)
		} else {
			columnTb(app, columns, db, tableName, dbName)
		}
	})
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Rune() == 'g' {
			s.ShowTables(app, db, dbName)
		}
		return event
	})
	form.SetTitle(dbName + "/?").SetBorder(true).SetBorderColor(tcell.ColorGreen)

	flex := tview.NewFlex()
	flex.AddItem(form, 0, 12, true)
	flex.AddItem(Info("\n\n Press ⎋esc or 'g' to go back\n\n"), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func columnTb(app *tview.Application, option string, db *gorm.DB, tbName, dbName string) {

	form := tview.NewForm()
	num, _ := strconv.Atoi(option)

	if num > 1 && num < 25 {
		for n := 0; n < num; n++ {
			form.AddInputField(fmt.Sprintf("%d column name", n+1), "", 20, nil, nil)
		}
	}

	form.AddButton("Create", func() {
		values := getQuery(form, num)
		mq.DbCreateTable(db, tbName, values)
		s.ShowDbs(app, db)
	})
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Rune() == 'g' {
			s.CreateTb(app, db, dbName)
		}
		return event
	})

	form.SetTitle(dbName + "/" + tbName + "/?").SetBorder(true).SetBorderColor(tcell.ColorGreen)

	flex := tview.NewFlex()
	flex.AddItem(form, 0, 12, true)
	flex.AddItem(Info("\n\n Press ⎋esc or 'g' to go back\n\n"), 0, 2, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func getQuery(form *tview.Form, num int) []string {
	var values []string

	for i := 0; i < num; i++ {
		if i == 0 {
			values = append(values, form.GetFormItem(i).(*tview.InputField).GetText()+"primary key not null ")
		} else {
			values = append(values, form.GetFormItem(i).(*tview.InputField).GetText()+" varchar(100) not null")
		}
	}
	return values
}

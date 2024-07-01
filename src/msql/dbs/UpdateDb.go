package dbs

import (
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/rivo/tview"
	"gorm.io/gorm"
	"log"
)

func (c CRUD) UpdateValue(app *tview.Application, db *gorm.DB, dbName, table, value string) {
	show := ShowS{}
	form := tview.NewForm()
	form.AddInputField(fmt.Sprintf("Column Name"),
		"", 30, nil, nil)
	form.AddInputField(fmt.Sprintf("Updating from %s.%s value=%s to value=", dbName, table, value),
		"", 30, nil, nil)

	form.AddButton("<-", func() { show.ShowValues(app, db, table, dbName) })
	form.AddButton("Update", func() {
		var err error
		var newValue string
		var columnName string
		app, db, newValue, columnName = fetchInputs(form, app, db, dbName, table, value)
		if err != nil {
			log.Fatal(err)
		}

		if newValue == "" || columnName == "" {
			form.GetFormItemByLabel("Column Name").(*tview.InputField).SetText("")
			form.GetFormItemByLabel(fmt.Sprintf("Updating from %s.%s value=%s to value=", dbName, table, value)).(*tview.InputField).SetText("")
			c.UpdateValue(app, db, dbName, table, value)
		} else {
			submit(app, columnName, newValue, db, dbName, table, value)
		}

	})

	form.SetBorder(true).SetTitle(fmt.Sprintf("Updating %s", dbName+"/"+table+"/"+table))

	if err := app.SetRoot(form, true).SetFocus(form).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func update(db *gorm.DB, tableName, columnName, valueBefore, valueAfter string) error {
	return mq.DbUpdate(db,
		tableName,
		valueAfter,
		columnName,
		valueBefore,
	)
}

func fetchInputs(
	form *tview.Form, app *tview.Application, db *gorm.DB, dbName string, table string, value string) (
	*tview.Application, *gorm.DB, string, string) {

	columnName := ""
	updatedValue := ""
	myColumns := mq.DbTableColumns(db, dbName, table)

	column := form.GetFormItem(0)
	updated := form.GetFormItem(1)

	if column != nil {
		columnName = column.(*tview.InputField).GetText()
	}
	if updated != nil {
		updatedValue = updated.(*tview.InputField).GetText()
	}
	if columnName == "" || updatedValue == "" || !ContainsString(myColumns, columnName) {
		return app, db, "", ""

	} else {
		if err := update(db, table, columnName, value, updatedValue); err != nil {
			return app, db, "", ""
		}

	}
	return app, db, updatedValue, columnName
}

func submit(app *tview.Application, columnName, newValue string, db *gorm.DB, dbName, table, value string) {
	cr := CRUD{}
	show := ShowS{}
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Do you want to change %s=%s", columnName, newValue)).
		AddButtons([]string{"Cancel", "Update"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Update" {
				show.ShowValues(app, db, table, dbName)
			} else {
				cr.UpdateValue(app, db, dbName, table, value)
			}
		})
	if err := app.SetRoot(modal, false).SetFocus(modal).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

package dbs

import (
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

type showI interface {
	ShowDbs(app *tview.Application, db *gorm.DB)
	ShowTables(app *tview.Application, db *gorm.DB, dbName string)
	ShowValues(app *tview.Application, db *gorm.DB, table string, dbName string)
}

type crud interface {
	CreateDb(app *tview.Application, db *gorm.DB) *tview.Form
	DeleteForm(app *tview.Application, db *gorm.DB) *tview.Form
	UpdateValue(app *tview.Application, db *gorm.DB, dbName, table, value string)
	CaptureDbName(form *tview.Form, db *gorm.DB, from int) string
	Layout(app *tview.Application, db *gorm.DB, dbName string, action string, from int)
}

type ShowS struct {
}
type CRUD struct {
}

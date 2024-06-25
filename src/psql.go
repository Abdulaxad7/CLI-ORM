package src

import (
	"Cli-Orm/config"
	"Cli-Orm/config/pq"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

var db *gorm.DB

func Psql(app *tview.Application) {

	form := tview.NewForm()
	form.AddButton("<-", func() { DbSelect(app) })
	form.AddInputField("Username", "", 20, nil, nil)
	form.AddInputField("Port", "", 20, nil, nil)
	form.AddInputField("Database", "", 20, nil, nil)

	form.AddPasswordField("Password", "", 20, '•', nil)
	form.SetBorder(true).SetTitle("Log in Postgresql").SetTitleAlign(tview.AlignCenter)

	form.AddButton("Submit", func() {
		app, db = checkPsql(form, app)

	})

	app.SetRoot(form, true).SetFocus(form).EnableMouse(true)
	return
}

func checkPsql(form *tview.Form, app *tview.Application) (*tview.Application, *gorm.DB) {
	var username string
	var port string
	var password string
	var dbName string
	usernameItem := form.GetFormItem(0)
	portItem := form.GetFormItem(1)
	dbN := form.GetFormItem(2)
	passwordItem := form.GetFormItem(3)

	if usernameItem != nil {
		username = usernameItem.(*tview.InputField).GetText()
	}
	if portItem != nil {
		port = portItem.(*tview.InputField).GetText()
	}
	if dbN != nil {
		dbName = passwordItem.(*tview.InputField).GetText()
	}
	if passwordItem != nil {
		password = passwordItem.(*tview.InputField).GetText()
	}
	if username == "" || port == "" || password == "" || dbName == "" {
		DbSelect(app)
	}

	db, err := pq.Connect(&config.DB{
		DBUser:     username,
		DBName:     dbName,
		DBPassword: password,
		Port:       port,
	})

	if err != nil {
		DbSelect(app)

	}

	return app, db
}

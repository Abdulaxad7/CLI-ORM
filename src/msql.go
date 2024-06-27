package src

import (
	"Cli-Orm/config"
	"Cli-Orm/config/mq"
	"Cli-Orm/src/msql"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func Msql(app *tview.Application) {
	var err error
	form := tview.NewForm()

	form.AddInputField("Username", "", 20, nil, nil)
	form.AddInputField("Port", "", 20, nil, nil)
	form.AddPasswordField("Password", "", 20, '•', nil)
	form.SetBorder(true).SetTitle("Log in Mysql").SetTitleAlign(tview.AlignCenter).SetBorderColor(tcell.ColorGreen)
	form.AddButton("Submit", func() {
		app, db, err = checkMsql(form, app)
		if err != nil {
			app.Stop()
			app1 := tview.NewApplication()
			DbSelect(app1)
		}
		msql.ShowDbs(app, db)

	})
	form.AddButton("<-", func() { DbSelect(app) })

	if err = app.SetRoot(form, true).SetFocus(form).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func checkMsql(form *tview.Form, app *tview.Application) (*tview.Application, *gorm.DB, error) {
	username := ""
	port := ""
	password := ""
	usernameItem := form.GetFormItem(0)
	portItem := form.GetFormItem(1)
	passwordItem := form.GetFormItem(2)

	if usernameItem != nil {
		username = usernameItem.(*tview.InputField).GetText()
	}
	if portItem != nil {
		port = portItem.(*tview.InputField).GetText()
	}

	if passwordItem != nil {
		password = passwordItem.(*tview.InputField).GetText()
	}
	if username == "" || port == "" || password == "" {
		DbSelect(app)
	}
	db, err := mq.Connect(&config.DB{
		DBUser:     username,
		DBPassword: password,
		Port:       port,
	})
	if err != nil {
		return app, db, err

	}

	return app, db, nil
}

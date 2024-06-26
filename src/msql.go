package src

import (
	"Cli-Orm/config"
	"Cli-Orm/config/mq"
	"Cli-Orm/src/msql"
	"fmt"
	"github.com/rivo/tview"
	"gorm.io/gorm"
	"log"
)

func Msql(app *tview.Application) {
	var err error
	form := tview.NewForm()
	form.AddButton("<-", func() { DbSelect(app) })
	form.AddInputField("Username", "", 20, nil, nil)
	form.AddInputField("Port", "", 20, nil, nil)
	form.AddPasswordField("Password", "", 20, '•', nil)
	form.SetBorder(true).SetTitle("Log in Mysql").SetTitleAlign(tview.AlignCenter)
	form.AddButton("Submit", func() {
		app, db, err = checkMsql(form, app)
		if err != nil {
			log.Fatal(err)
		}
		msql.ShowDbs(app, db)

	})

	app.SetRoot(form, true).SetFocus(form).EnableMouse(true)

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
		fmt.Errorf("username or port or password can't be empty")
		DbSelect(app)
	}

	db, err := mq.Connect(&config.DB{
		DBUser:     username,
		DBPassword: password,
		Port:       port,
	})

	if err != nil {
		DbSelect(app)
	}

	return app, db, nil
}

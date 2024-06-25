package screen

import (
	"Cli-Orm/config"
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/gdamore/tcell/v2"
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
		ShowDbs(app, db)

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
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return app, db, nil
}

func ShowDbs(app *tview.Application, db *gorm.DB) {
	tb := tview.NewTable()
	tb.SetBorders(true).SetBorder(true)
	var results []map[string]interface{}
	db.Raw("SHOW DATABASES;").Scan(&results)
	for i, v := range results {
		for _, v2 := range v {
			tb.SetCell(i, 0, &tview.TableCell{Text: v2.(string), Align: tview.AlignCenter, Color: tview.Styles.TitleColor}).SetSelectable(true, false).SetOffset(1, 1)
			tb.Select(i, 0).SetSelectedFunc(func(row int, column int) {
				tb.GetCell(row, column).SetTextColor(tcell.ColorRed)
				tb.SetSelectable(false, false).ScrollToBeginning()
			})
		}
	}
	app.SetRoot(tb, true).SetFocus(tb)

}

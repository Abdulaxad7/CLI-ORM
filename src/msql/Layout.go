package msql

import (
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
	"strings"
)

func (c CRUD) Layout(app *tview.Application, db *gorm.DB, dbName string, action string, from int) {
	show := ShowS{}
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Are you sure you want to %s %s", action, dbName))
	act := strings.ToUpper(string(action[0])) + action[1:]
	modal.AddButtons([]string{"Cancel", act}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				show.ShowDbs(app, db)
			} else {
				if from == 0 {
					mq.DbCreate(db, dbName)
				} else {
					mq.DbDrop(db, dbName)
				}
				show.ShowDbs(app, db)
			}
		})
	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			show.ShowDbs(app, db)
		}
		return event
	})
	if err := app.SetRoot(modal, true).SetFocus(modal).Run(); err != nil {
		panic(err)
	}
}

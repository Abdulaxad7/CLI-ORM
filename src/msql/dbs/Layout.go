package dbs

import (
	"Cli-Orm/config/mq"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gorm.io/gorm"
	"strings"
	"unicode"
)

func (c CRUD) Layout(app *tview.Application, db *gorm.DB, dbName, tableName, action string, from int) {
	show := ShowS{}
	modal := tview.NewModal().
		SetText(fmt.Sprintf("Are you sure you want to %s %s.%s", action, dbName, tableName))
	act := strings.ToUpper(string(action[0])) + action[1:]
	modal.AddButtons([]string{"Cancel", act}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				show.ShowDbs(app, db)
			} else {
				if from == 0 {
					mq.DbCreate(db, dbName)
				} else if from == 1 {
					mq.DbDrop(db, dbName)
				} else if from == 2 {
					mq.DbDropTable(db, dbName, tableName)
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

func Info(setText string) *tview.TextView {
	text := tview.NewTextView()
	text.SetText(setText)
	text.SetBorderColor(tcell.ColorBlack)
	text.SetTextColor(tcell.ColorGrey)
	return text
}

func ContainsAlpha(s string) bool {
	for _, n := range s {
		if unicode.IsLetter(n) {
			return true
		}
	}
	return false
}

func ContainsString(results []map[string]interface{}, target string) bool {
	for _, m := range results {
		for _, v := range m {
			if str, ok := v.(string); ok && str == target {
				return true
			}
		}
	}
	return false
}
func CutMapStringInterface(data []map[string]interface{}) []string {
	var result []string
	for _, v1 := range data {
		for _, v2 := range v1 {
			result = append(result, v2.(string))
		}
	}
	return result
}

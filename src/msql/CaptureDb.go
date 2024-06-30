package msql

import (
	"Cli-Orm/config/mq"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

func (c CRUD) CaptureDbName(form *tview.Form, db *gorm.DB, from int) string {
	selectAllDb := mq.DbQuery(db)
	dbN := ""
	dbName := form.GetFormItem(0)
	if dbName != nil {
		dbN = dbName.(*tview.InputField).GetText()
	}
	if from == 0 {
		if !ContainsString(selectAllDb, dbN) {
			return dbN
		}
	} else {
		if ContainsString(selectAllDb, dbN) {
			return dbN
		}
	}
	return ""
}

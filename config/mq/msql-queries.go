package mq

import (
	"fmt"
	"github.com/rivo/tview"
	"gorm.io/gorm"
	"os"
)

func DbQuery(db *gorm.DB) []map[string]interface{} {
	var results []map[string]interface{}
	raws := db.Raw("SHOW DATABASES;").Scan(&results)
	checkError(raws)
	return results
}
func DbUseQuery(db *gorm.DB, dbName string) {
	tx := db.Exec(fmt.Sprintf("USE %s", dbName))
	checkError(tx)
}
func DbTables(db *gorm.DB) []map[string]interface{} {
	var results []map[string]interface{}
	raws := db.Raw("SHOW TABLES;").Scan(&results)
	checkError(raws)
	return results
}
func DbValues(db *gorm.DB, table string) []map[string]interface{} {
	var results []map[string]interface{}
	raws := db.Raw(fmt.Sprintf("SELECT * FROM %s;", table)).Find(&results)
	checkError(raws)
	return results
}
func DbUpdate(db *gorm.DB, tableName, valueAfter, columnName, valueBefore string) error {
	db.Exec("SET SQL_SAFE_UPDATES = 0")
	tx := db.Exec(fmt.Sprintf("UPDATE %s set %s=%s where %s=%s", tableName, columnName, valueAfter, columnName, valueBefore))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func DbTableColumns(db *gorm.DB, dbName, table string) []map[string]interface{} {
	var myColumns []map[string]interface{}
	query := fmt.Sprintf("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA='%s'  AND  TABLE_NAME='%s' ;", dbName, table)
	raws := db.Raw(query).Scan(&myColumns)
	checkError(raws)
	return myColumns
}

func checkError(err *gorm.DB) {
	if err.Error != nil {
		app := tview.NewApplication()
		pages := tview.NewPages()
		for page := 0; page < 5; page++ {
			func(page int) {
				pages.AddPage(fmt.Sprintf("page-%d", page),
					tview.NewModal().
						SetText(fmt.Sprintf("Something went wrong in database!!!\n%s", err.Error)).
						AddButtons([]string{"Quit"}).
						SetDoneFunc(func(buttonIndex int, buttonLabel string) {
							app.Stop()
							os.Exit(0)
						}),
					false,
					page == 0)
			}(page)
		}
		if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
			panic(err)
		}
	}

}

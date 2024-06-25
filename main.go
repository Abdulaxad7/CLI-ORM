package main

import (
	"Cli-Orm/screen"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/rivo/tview"
)

func main() {
	//db, err := pq.Connect(&config.DB{
	//	DBUser:     "root",
	//	DBPassword: "0987poiulkjh",
	//	DBAddr:     "3306",
	//	DBName:     "postgres",
	//})
	//
	//d, err := db.Query("select *from teacher_dbs")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(d)
	app := tview.NewApplication()
	screen.DbSelect(app)

}

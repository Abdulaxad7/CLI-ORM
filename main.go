package main

import (
	"Cli-Orm/src"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	src.DbSelect(app)
}

package main

import (
	"github.com/rakutentech/code-coverage-dashboard/app"
	"github.com/rakutentech/code-coverage-dashboard/models"
)

func main() {
	app.SetEnv()    // load .env
	app.SetAppLog() // set application log based on .env and then config
	db := app.NewDB()
	err := db.AutoMigrate(&models.Coverage{})
	if err != nil {
		panic(err)
	}
}

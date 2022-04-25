package main

import (
	"github.com/k0kubun/pp"
	"github.com/rakutentech/code-coverage-dashboard/app"
	"github.com/rakutentech/code-coverage-dashboard/config"
)

func main() {
	app.SetEnv()    // load .env
	app.SetAppLog() // set application log based on .env and then config
	conf := config.NewConfig()
	//#nosec
	pp.Println(conf)
}

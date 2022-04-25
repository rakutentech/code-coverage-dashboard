package app

import (
	"fmt"
	"log"

	"github.com/rakutentech/code-coverage-dashboard/config"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// SetAppLog sets the logger format and log rotation
func SetAppLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(Logger("app.log"))
}

// Logger returns a new logger with log rotation
// Is also used as http middleware for logging http requests of echo
func Logger(logFilename string) *lumberjack.Logger {
	conf := config.NewConfig()
	fn := conf.AppConfig.LogDir + "/" + logFilename
	fmt.Println("Logs file" + fn)
	return &lumberjack.Logger{
		Filename:   fn,
		MaxSize:    50, // megabytes
		MaxBackups: 10,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default

	}
}

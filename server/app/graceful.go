package app

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo/v4"
	"github.com/rakutentech/code-coverage-dashboard/config"
)

// GracefulServerWithPid reloads server with pid
// kill -HUP when binary is changed
// kill -9 when want to kill the process and make the application dead and want to restart
// kill -9 is NOT FOR FAINT HEARTED and must not be done on prod unless SOUT
func GracefulServerWithPid(e *echo.Echo, port string) {
	conf := config.NewConfig()
	log.Print(pp.Sprint(conf))
	e.Server.Addr = conf.AppConfig.AppHost + ":" + port
	server := endless.NewServer(e.Server.Addr, e)
	server.BeforeBegin = func(add string) {
		log.Print("info: actual pid is", syscall.Getpid())
		pidFile := filepath.Join(conf.AppConfig.PidDir, port+".pid")
		err := os.Remove(pidFile)
		if err != nil {
			log.Print("error: pid file error: ", err)
		} else {
			log.Print("success: pid file success", pidFile)
		}
		err = ioutil.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0600)
		if err != nil {
			log.Print("error: write pid file error: ", err)
		} else {
			log.Print("success: write pid file success", pidFile)
		}
	}
	if err := server.ListenAndServe(); err != nil {
		log.Print("critical: graceful error: ", err)
	}
}

package main

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/rakutentech/code-coverage-dashboard/app"
	"github.com/rakutentech/code-coverage-dashboard/config"
	"github.com/rakutentech/code-coverage-dashboard/middlewares"
	"github.com/rakutentech/code-coverage-dashboard/route"
)

func main() {
	createApplication()
	conf := config.NewConfig() // set config once from .env
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(...) are middlewares globally for this application

	// Capabilities
	// 1. Handle http logs
	// 2. Handle global http errors and notify alerts
	middlewares.HTTPMiddleware(e)
	e.HTTPErrorHandler = middlewares.HTTPErrorHandler

	// Capabilities:
	// 1. Able to recover when connect to DB fails
	e.Use(middleware.Recover())

	if conf.AppConfig.AppEnv == "local" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3009"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	}

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(conf.AppConfig.SessionSecret))))

	// All routes for this application
	route.RegisterRoutes(e)

	port := "3001"
	if len(os.Args) > 2 {
		port = os.Args[1]
	}

	app.GracefulServerWithPid(e, port)
}

// createApplication creates the application
// and initializes the envs and configuration
// and the log
func createApplication() {
	app.SetEnv()    // load .env
	app.SetAppLog() // set application log based on .env and then config
}

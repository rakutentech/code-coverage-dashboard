package middlewares

import (
	"fmt"
	"log"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/rakutentech/code-coverage-dashboard/config"
)

// UserCanReadRepo is a middleware to allow access to fail server
func UserCanReadRepo() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		conf := config.NewConfig()
		_ = conf
		return func(c echo.Context) error {
			log.Println(c.Request().URL.Path)
			path := strings.Split(c.Request().URL.Path, "/")

			if len(path) >= 0 && path[0] != "" {
				return fmt.Errorf("invalid URL")
			}
			// buggy and untested. Purpose is to restrict access to file server assets
			// by making a Github api call to check user permissions for the repo

			// append / to app brase url since split would have stripped it
			// if len(path) >= 1 && "/"+path[1] != conf.AppConfig.AppBaseURL {
			// 	return fmt.Errorf("invalid URL, expected %s", conf.AppConfig.AppBaseURL)
			// }
			// if len(path) >= 2 && path[2] != "assets" {
			// 	return fmt.Errorf("invalid URL, must be assets")
			// }
			// # TDOO enable or disable restrictions
			// if len(path) >= 3 && path[3] == "" {
			// 	return fmt.Errorf("invalid URL, must have org name")
			// }
			// if len(path) >= 4 && path[4] == "" {
			// 	return fmt.Errorf("invalid URL, must have branch name")
			// }

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}

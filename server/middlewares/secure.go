package middlewares

import (
	echo "github.com/labstack/echo/v4"
	"github.com/rakutentech/code-coverage-dashboard/config"
)

// SecureMiddleware is a middleware that sets URL.Schema to https
// By doing so, the session handler sets cookie secure by bypassing tls verification for reverse proxy
func SecureMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		conf := config.NewConfig()
		return func(c echo.Context) error {
			if conf.AppConfig.AppSecure {
				c.Request().URL.Scheme = "https"
			}
			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}

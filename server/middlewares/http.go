package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/rakutentech/code-coverage-dashboard/app"

	echo "github.com/labstack/echo/v4"
	midw "github.com/labstack/echo/v4/middleware"
)

// HTTPErrorResponse is the response for HTTP errors
type HTTPErrorResponse struct {
	Error interface{} `json:"error"`
}

// HTTPMiddleware global middleware for all requests
func HTTPMiddleware(e *echo.Echo) {
	e.Pre(midw.RemoveTrailingSlash())

	e.Use(midw.LoggerWithConfig(midw.LoggerConfig{
		Format: time.Now().Format("2006/01/02 15:04:05") + ` {"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			// `"header":"${header:Authorization}"` + // for debug of Github token
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		Output: app.Logger("app.log"),
	}))
}

// HTTPErrorHandler handles HTTP errors for entire application
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		message = err.Error()
	}

	if err = c.JSON(code, &HTTPErrorResponse{Error: message}); err != nil {
		log.Print("error: ", err)
	}
}

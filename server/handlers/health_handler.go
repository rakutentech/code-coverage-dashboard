package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/rakutentech/code-coverage-dashboard/config"
)

// HealthCheckRequest for the /health
type HealthCheckRequest struct {
}

// HealthCheckResponse for /auth/health check
// returns success if server is healthy
type HealthCheckResponse struct {
	Success string `json:"success"`
}

// HealthHandler injected with dependendencies to be used in the controller
// Passed to routes and defined methods
type HealthHandler struct {
}

// NewHealthHandler returns a new HealthHandler interface
func NewHealthHandler() HealthHandler {
	return HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	conf := config.NewConfig()
	request := &HealthCheckRequest{}
	response := &HealthCheckResponse{}
	err := c.Bind(request)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response.Success = conf.AppConfig.AppEnv
	return c.JSON(http.StatusOK, response)
}

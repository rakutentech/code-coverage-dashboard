package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	Setup()
}

func TestHealthController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/coverages-api/health", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewHealthHandler()

	// Assertions
	err := h.HealthCheck(c)
	if err != nil {
		he, _ := err.(*echo.HTTPError)
		assert.Equal(t, http.StatusOK, he.Code)
	} else {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

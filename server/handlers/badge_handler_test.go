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
func TestBadgeShow(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		wantStatusCode int
	}{
		{
			name:           "Test coverage upload",
			url:            "/badge",
			method:         http.MethodGet,
			wantStatusCode: http.StatusOK,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(test.method, test.url, nil)

			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := NewBadgeHandler()
			err := h.BadgeShow(ctx)
			if err != nil {
				he, _ := err.(*echo.HTTPError)
				assert.Equal(t, test.wantStatusCode, he.Code)
			} else {
				assert.Equal(t, test.wantStatusCode, rec.Code)
			}

		})
	}
}

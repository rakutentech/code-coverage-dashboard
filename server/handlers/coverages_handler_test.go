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
func TestCoveragesApiValidations(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		method         string
		wantStatusCode int
	}{
		{
			name:           "Test coverage upload",
			url:            "/coverages-api",
			method:         http.MethodPost,
			wantStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:           "Test coverage sync",
			url:            "/coverages-api",
			method:         http.MethodPut,
			wantStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:           "Test coverage upload",
			url:            "/coverages-api/badge",
			method:         http.MethodGet,
			wantStatusCode: http.StatusUnprocessableEntity,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(test.method, test.url, nil)
			req.Header.Set("Content-Type", "multipart/form-data")

			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := NewCoveragesHandler()
			err := h.CoveragesUpload(ctx)
			if err != nil {
				he, _ := err.(*echo.HTTPError)
				assert.Equal(t, test.wantStatusCode, he.Code)
			} else {
				assert.Equal(t, test.wantStatusCode, rec.Code)
			}

		})
	}
}

package handlers

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	echo "github.com/labstack/echo/v4"
	badge "github.com/narqo/go-badge"
	"github.com/rakutentech/code-coverage-dashboard/models/repositories"
	"github.com/rakutentech/code-coverage-dashboard/services"
)

// BadgeHandler injected with dependendencies to be used in the controller
// Passed to routes and defined methods
type BadgeHandler struct {
	coveragesRepository *repositories.CoveragesRepository
}

// NewBadgeHandler returns a new BadgeHandler interface
func NewBadgeHandler() BadgeHandler {
	return BadgeHandler{
		coveragesRepository: repositories.NewCoveragesRepository(),
	}
}

func (h *BadgeHandler) BadgeShow(c echo.Context) error {
	request := &BadgeRequest{}
	title := ""
	subtitle := "Invalid"

	err := c.Bind(request)
	if err != nil {
		return streamBadge(c, title, subtitle, badge.ColorRed)
	}

	// validate request
	_, err = services.ValidateRequest(request)
	if err != nil {
		log.Print("Validation Errors: ", err.Error())
		return streamBadge(c, title, subtitle, badge.ColorRed)
	}

	title = request.BranchName
	if request.Subtitle != "" {
		subtitle = request.Subtitle
		return streamBadge(c, title, subtitle, badge.ColorBrightgreen)
	}
	coverage, err := h.coveragesRepository.FindCoverage(request.OrgName, request.RepoName, request.BranchName, request.Language)
	if err != nil {
		log.Print("ERROR: ", err.Error())
		return streamBadge(c, title, subtitle, badge.ColorRed)
	}

	// float to string
	subtitle = strconv.FormatFloat(float64(coverage.Percentage), 'f', 2, 64) + "%"
	return streamBadge(c, title, subtitle, badge.ColorGreen)

}

func streamBadge(c echo.Context, title, subtitle string, color badge.Color) error {
	// process badge
	var buf bytes.Buffer

	err := badge.Render(title, subtitle, color, &buf)

	if err != nil {
		log.Print("ERROR: ", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	tenMinsBefore := time.Now().UTC().Add((-10) * time.Minute).Format("Mon, 02 Jan 2006 15:04:05 GMT")

	c.Response().Header().Set("Cache-Control", "no-cache,max-age=0")
	c.Response().Header().Set("Expires", tenMinsBefore)
	return c.Stream(http.StatusOK, "image/svg+xml", &buf)
}

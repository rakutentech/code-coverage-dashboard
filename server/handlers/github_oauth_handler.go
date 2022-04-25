package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
	echo "github.com/labstack/echo/v4"
	"github.com/rakutentech/code-coverage-dashboard/config"
	"github.com/rakutentech/code-coverage-dashboard/services"
)

type LoginGithubCallbackRequest struct {
	Code string `form:"code" query:"code" json:"code"  validate:"required" message:"code is required"`
}

// GithubOAuthHandler injected with dependendencies to be used in the controller
// Passed to routes and defined methods
type GithubOAuthHandler struct {
	githubOAuthService *services.GithubOAuthService
	sessionService     *services.SessionService
}

// NewGithubOAuthHandler returns a new GithubOAuthHandler interface
func NewGithubOAuthHandler() GithubOAuthHandler {
	return GithubOAuthHandler{
		githubOAuthService: services.NewGithubOAuthService(),
		sessionService:     services.NewSessionService(),
	}
}

func (h *GithubOAuthHandler) LoginGithub(c echo.Context) error {
	conf := config.NewConfig()
	var protocol string
	protocol = "http://"
	if conf.AppConfig.AppSecure {
		protocol = "https://"
	}

	// Create the dynamic redirect URL for login
	redirectURL := fmt.Sprintf(
		"%s/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		conf.GithubConfig.GithubURL,
		conf.GithubConfig.ClientID,
		protocol+conf.AppConfig.AppURL+conf.AppConfig.AppBaseURL+"/login/github/callback",
	)
	log.Println("redirectURL:", redirectURL)
	// redirect to login
	// return c.JSON(http.StatusOK, "123")
	return c.Redirect(http.StatusMovedPermanently, redirectURL)
}

func (h *GithubOAuthHandler) LoginGithubCallback(c echo.Context) error {
	request := &LoginGithubCallbackRequest{}
	err := c.Bind(request)
	log.Print(pp.Sprint(request))
	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	githubAccessToken, err := h.githubOAuthService.GetGithubAccessToken(request.Code)

	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// update session
	err = h.sessionService.Set(c, services.GithubAccessTokenSessionKey, githubAccessToken)
	log.Print("githubAccessToken: ", githubAccessToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	githubData, err := h.githubOAuthService.GetGithubUser(githubAccessToken)
	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, githubData)
}

func (h *GithubOAuthHandler) LoginGithubStatus(c echo.Context) error {
	sess, err := h.sessionService.Get(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	githubAccessToken := sess.Values[services.GithubAccessTokenSessionKey]
	if githubAccessToken == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Not logged in")
	}
	githubData, err := h.githubOAuthService.GetGithubUser(githubAccessToken.(string))
	if err != nil {
		log.Print("error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, githubData)
}

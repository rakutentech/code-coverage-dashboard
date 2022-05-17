package route

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakutentech/code-coverage-dashboard/config"
	"github.com/rakutentech/code-coverage-dashboard/handlers"
	m "github.com/rakutentech/code-coverage-dashboard/middlewares"
)

// RegisterRoutes register all routes
func RegisterRoutes(e *echo.Echo) {

	conf := config.NewConfig()
	healthHandler := handlers.NewHealthHandler()
	e.GET(
		conf.AppConfig.AppBaseURL+"/health",
		healthHandler.HealthCheck,
		m.SecureMiddleware(),
	)

	badgeHandler := handlers.NewBadgeHandler()
	e.GET(
		conf.AppConfig.AppBaseURL+"/badge",
		badgeHandler.BadgeShow,
		m.SecureMiddleware(),
	)

	coveragesHandler := handlers.NewCoveragesHandler()
	e.GET(conf.AppConfig.AppBaseURL, coveragesHandler.CoveragesPaginate, m.SecureMiddleware())
	e.PUT(conf.AppConfig.AppBaseURL, coveragesHandler.CoveragesUpdateBranches, m.SecureMiddleware())
	e.POST(conf.AppConfig.AppBaseURL, coveragesHandler.CoveragesUpload, m.SecureMiddleware())

	githubOAuthHandler := handlers.NewGithubOAuthHandler()
	// Login route
	e.GET(conf.AppConfig.AppBaseURL+"/login/github", githubOAuthHandler.LoginGithub)

	// Github callback
	e.GET(conf.AppConfig.AppBaseURL+"/login/github/callback", githubOAuthHandler.LoginGithubCallback)
	// Github check if user is logged in
	e.GET(conf.AppConfig.AppBaseURL+"/login/github/status", githubOAuthHandler.LoginGithubStatus)

	// coverages-api/assets/
	fileServerReverseProxy(e)
}

func fileServerReverseProxy(e *echo.Echo) {
	conf := config.NewConfig()
	assetsURI := conf.AppConfig.AppBaseURL + "/assets"
	g := e.Group(assetsURI)

	// Setup reverse proxy
	fileServerURL, err := url.Parse(conf.AppConfig.FileServerURL)
	if err != nil {
		panic(err) // should never happen
	}
	g.Use(m.SecureMiddleware())
	g.Use(m.UserCanReadRepo())
	g.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Skipper: func(e echo.Context) bool {
			if e.Request().Method != http.MethodGet {
				// skip non GET requests
				return true
			}
			if strings.HasPrefix(e.Request().URL.Path, assetsURI) {
				if strings.HasPrefix(e.Request().URL.Path, assetsURI+"/") {
					e.Request().URL.Path = strings.Replace(e.Request().URL.Path, assetsURI, "", 1)
					ext := filepath.Ext(e.Request().URL.Path)
					if ext == "" {
						e.Request().URL.Path = e.Request().URL.Path + "/" // append / to the end
					}
				} else {
					e.Request().URL.Path = strings.Replace(e.Request().URL.Path, assetsURI, "", 1)
				}
			}
			// ok, donot skip this, legit assets reverse proxy call
			return false
		},
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{URL: fileServerURL}}),
	}))
}

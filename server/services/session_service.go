package services

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	echo "github.com/labstack/echo/v4"
)

type SessionService struct {
	GithubAccessTokenSessionKey string
}

const (
	// #nosec
	GithubAccessTokenSessionKey = "ccd_github_access_token"
)

// NewSessionService creates a new SessionService
func NewSessionService() *SessionService {
	return &SessionService{
		GithubAccessTokenSessionKey: GithubAccessTokenSessionKey,
	}
}

func (s *SessionService) Get(c echo.Context) (*sessions.Session, error) {
	sess, err := session.Get(GithubAccessTokenSessionKey, c)
	return sess, err
}
func (s *SessionService) Set(c echo.Context, key, value string) error {
	sess, err := session.Get(key, c)
	if err != nil {
		return err
	}
	sess.Values[key] = value
	err = sess.Save(c.Request(), c.Response())
	return err
}

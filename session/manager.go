package session

import (
	"errors"

	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
)

var GlobalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "sid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go GlobalSessions.GC()
}

func SetCurrentUser(c web.Controller, u *models.User) error {
	err := c.SetSession("username", u.Username)
	if err != nil {
		return err
	}
	err = c.SetSession("password", u.Password)
	if err != nil {
		return err
	}
	err = c.SetSession("user_obj", u)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentUser(c web.Controller) (*models.User, error) {
	u, ok := c.GetSession("user_obj").(*models.User)
	if !ok {
		return nil, errors.New("invalid user in session")
	}
	return u, nil
}

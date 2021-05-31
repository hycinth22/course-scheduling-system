package session

import (
	"errors"
	"log"

	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
)

//var GlobalSessions *session.Manager

func init() {
	//sessionConfig := &session.ManagerConfig{
	//	CookieName:      "sid",
	//	EnableSetCookie: true,
	//	Gclifetime:      3600,
	//	Maxlifetime:     3600,
	//	Secure:          false,
	//	CookieLifeTime:  3600,
	//	ProviderConfig:  "./tmp",
	//}
	//GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	//go GlobalSessions.GC()
}

func SetCurrentUser(c *web.Controller, u *models.User) error {
	log.Println("SetCurrentUser", *u)
	err := c.SetSession("username", u.Username)
	if err != nil {
		log.Println("SetCurrentUser", err)
		return err
	}
	err = c.SetSession("password", u.Password)
	if err != nil {
		log.Println("SetCurrentUser", err)
		return err
	}
	err = c.SetSession("user_obj", *u)
	if err != nil {
		log.Println("SetCurrentUser", err)
		return err
	}
	return nil
}

func GetCurrentUser(c *web.Controller) (*models.User, error) {
	log.Println("GetCurrentUser", c.GetSession("username"), c.GetSession("password"))
	log.Println("GetCurrentUser", c.GetSession("user_obj"))
	obj := c.GetSession("user_obj")
	if obj == nil {
		return nil, errors.New("no current user")
	}
	u, ok := obj.(models.User)
	if !ok {
		return nil, errors.New("invalid user in session")
	}
	return &u, nil
}

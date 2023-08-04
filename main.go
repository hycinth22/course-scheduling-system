package main

import (
	"courseScheduling/models"
	"encoding/gob"

	_ "courseScheduling/routers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	gob.Register(models.User{})
}

func main() {
	// api docs
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// session
	web.BConfig.WebConfig.Session.SessionOn = true
	web.BConfig.WebConfig.Session.SessionProvider = "file"
	web.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	// launch now
	web.Run()
}

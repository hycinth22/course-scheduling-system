package main

import (
	"encoding/gob"

	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web/filter/cors"

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
	// allow cors
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	web.BConfig.WebConfig.Session.SessionOn = true
	// launch now
	web.Run()
}

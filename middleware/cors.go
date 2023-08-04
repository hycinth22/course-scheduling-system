package middleware

import "github.com/beego/beego/v2/server/web/filter/cors"

var CORS = cors.Allow(&cors.Options{
	AllowAllOrigins:  false,
	AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	AllowCredentials: true,
})

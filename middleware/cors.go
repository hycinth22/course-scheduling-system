package middleware

import "github.com/beego/beego/v2/server/web/filter/cors"

var CORS = cors.Allow(&cors.Options{
	AllowAllOrigins:  false,
	AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Authorization"},
	ExposeHeaders:    []string{"Content-length"},
	AllowCredentials: true,
})

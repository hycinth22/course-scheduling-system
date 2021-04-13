package routers

import (
	"courseScheduling/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	admin := web.NewNamespace("/admin",
		web.NSNamespace("/user", web.NSInclude(&controllers.UserController{})),
		web.NSNamespace("/course", web.NSInclude(&controllers.CourseController{})),
	)
	view := web.NewNamespace("/view")
	web.AddNamespace(admin, view)
}

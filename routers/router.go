// @APIVersion 1.0.0
// @Title course scheduling system - backend
// @Description course scheduling system - backend
// @Contact 206950850@qq.com
package routers

import (
	"courseScheduling/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	admin := web.NewNamespace("/admin",
		web.NSNamespace("/clazz", web.NSInclude(&controllers.ClazzController{})),
		web.NSNamespace("/college", web.NSInclude(&controllers.CollegeController{})),
		web.NSNamespace("/course", web.NSInclude(&controllers.CourseController{})),
		web.NSNamespace("/dept", web.NSInclude(&controllers.DepartmentController{})),
		web.NSNamespace("/schedule", web.NSInclude(&controllers.ScheduleController{})),
		web.NSNamespace("/semester", web.NSInclude(&controllers.SemesterController{})),
		web.NSNamespace("/timespan", web.NSInclude(&controllers.TimespanController{})),
		web.NSNamespace("/user", web.NSInclude(&controllers.UserController{})),
	)
	view := web.NewNamespace("/view")
	web.AddNamespace(admin, view)
}

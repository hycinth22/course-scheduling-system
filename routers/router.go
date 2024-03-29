// @APIVersion 1.0.0
// @Title course scheduling system - backend
// @Description course scheduling system - backend
// @Contact 206950850@qq.com
package routers

import (
	"courseScheduling/controllers"
	"courseScheduling/middleware"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	auth := web.NewNamespace("/api",
		web.NSNamespace("/backup", web.NSInclude(&controllers.BackupController{})),
		web.NSNamespace("/clazz", web.NSInclude(&controllers.ClazzController{})),
		web.NSNamespace("/clazzroom", web.NSInclude(&controllers.ClazzroomController{})),
		web.NSNamespace("/college", web.NSInclude(&controllers.CollegeController{})),
		web.NSNamespace("/course", web.NSInclude(&controllers.CourseController{})),
		web.NSNamespace("/dashboard", web.NSInclude(&controllers.DashboardController{})),
		web.NSNamespace("/dept", web.NSInclude(&controllers.DepartmentController{})),
		web.NSNamespace("/instruct", web.NSInclude(&controllers.InstructController{})),
		web.NSNamespace("/schedule", web.NSInclude(&controllers.ScheduleController{})),
		web.NSNamespace("/semester", web.NSInclude(&controllers.SemesterController{})),
		web.NSNamespace("/teacher", web.NSInclude(&controllers.TeacherController{})),
		web.NSNamespace("/timespan", web.NSInclude(&controllers.TimespanController{})),
	)
	nonauth := web.NewNamespace("/api",
		web.NSNamespace("/user", web.NSInclude(&controllers.UserController{})),
	)
	// allow cors
	auth.Filter("before", middleware.CORS)
	nonauth.Filter("before", middleware.Auth)
	// auth
	auth.Filter("before", middleware.Auth)
	view := web.NewNamespace("/view")
	web.AddNamespace(auth, nonauth, view)
}

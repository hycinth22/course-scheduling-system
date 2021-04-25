package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["courseScheduling/controllers:ClazzController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ClazzController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ClazzroomController"],
		beego.ControllerComments{
			Method:           "ImportFromExcel",
			Router:           "/excel",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:CollegeController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:CollegeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:cid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:cid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:cid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "ImportFromExcel",
			Router:           "/excel",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:DepartmentController"],
		beego.ControllerComments{
			Method:           "ImportFromExcel",
			Router:           "/excel",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "GetSchedule",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "DeleteAllScheduleInSemester",
			Router:           "/",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "DeleteSchedule",
			Router:           "/:schedule_id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "GetScheduleItems",
			Router:           "/:schedule_id/items",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "GetScheduleItemsGroupView",
			Router:           "/:schedule_id/items/group_view",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "ScheduleDownloadStudentExcel",
			Router:           "/:schedule_id/student_excel",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "ScheduleDownloadTeacherExcel",
			Router:           "/:schedule_id/teacher_excel",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:ScheduleController"],
		beego.ControllerComments{
			Method:           "NewSchedule",
			Router:           "/new",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:SemesterController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:SemesterController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:TimespanController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:TimespanController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseScheduling/controllers:UserController"] = append(beego.GlobalControllerRouter["courseScheduling/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

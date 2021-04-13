package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:cid",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:cid",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:cid",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"] = append(beego.GlobalControllerRouter["courseSchduling/controllers:CourseController"],
		beego.ControllerComments{
			Method:           "ImportFromExcel",
			Router:           "/excel",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["courseSchduling/controllers:UserController"] = append(beego.GlobalControllerRouter["courseSchduling/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

package controllers

import (
	"log"

	"courseScheduling/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = map[string]interface{}{
			"code": 0,
			"msg":  "login success",
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": -10001,
			"msg":  "user not exist",
		}
	}
	err := u.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

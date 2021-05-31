package controllers

import (
	"encoding/json"
	"log"
	"time"

	"courseScheduling/models"
	"courseScheduling/session"

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
func (this *UserController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	if ok, u := models.CanLogin(username, password); ok {
		this.Data["json"] = map[string]interface{}{
			"code": 0,
			"msg":  "login success",
			"profile": map[string]interface{}{
				"username":           u.Username,
				"role":               u.Role,
				"lastTime":           u.LastLogin,
				"lastLoc":            u.LastLoc,
				"associated_teacher": u.AssociatedTeacher.Id,
			},
		}
		go func() {
			err := models.UpdateLogin(u, time.Now(), getIPLoc(this.Ctx.Input.IP()))
			if err != nil {
				log.Println(err)
			}
		}()
		err := session.SetCurrentUser(&this.Controller, u)
		if err != nil {
			log.Println(err)
		}
	} else if !ok && u != nil && u.Status != 0 {
		this.Data["json"] = map[string]interface{}{
			"code": -10002,
			"msg":  "user banned",
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": -10001,
			"msg":  "user not exist",
		}
	}
	err := this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /list [get]
func (this *UserController) GetAll() {
	var query struct {
		Search    string `form:"search"`
		PageIndex int    `form:"pageIndex"`
		PageSize  int    `form:"pageSize"`
	}
	if err := this.ParseForm(&query); err != nil {
		log.Println(err)
	}
	var (
		courses []*models.User
		total   int
	)
	if query.Search == "" {
		courses, total = models.ListUsers(getOffset(query.PageIndex, query.PageSize), query.PageSize)
	} else {
		courses, total = models.SearchUsers(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Search)
	}
	this.Data["json"] = map[string]interface{}{
		"list":      courses,
		"pageTotal": total,
	}
	err := this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /login_state [get]
func (this *UserController) LoginState() {
	username, ok1 := this.GetSession("username").(string)
	password, ok2 := this.GetSession("password").(string)
	log.Println(username, password)
	if ok, u := models.CanLogin(username, password); ok1 && ok2 && ok {
		this.Data["json"] = map[string]interface{}{
			"code":     0,
			"username": username,
			"profile": map[string]interface{}{
				"username":           u.Username,
				"role":               u.Role,
				"lastTime":           u.LastLogin,
				"lastLoc":            u.LastLoc,
				"associated_teacher": u.AssociatedTeacher.Id,
			},
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code": -10001,
		}
	}
	err := this.ServeJSON()
	if err != nil {
		return
	}
}

// @router /:id/status [put]
func (this *UserController) SetStatus() {
	id, err := this.GetInt(":id")
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	var body struct {
		Value int
	}
	err = this.ParseForm(&body)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.UpdateUserStatus(&models.User{Id: id, Status: body.Value})
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
}

// @router /:id/password [put]
func (this *UserController) ResetPassword() {
	id, err := this.GetInt(":id")
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	var body struct {
		Value string
	}
	err = this.ParseForm(&body)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.UpdateUserPassword(&models.User{Id: id, Password: body.Value})
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
}

// @router /new [post]
func (this *UserController) Create() {
	var c models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.AddUser(c)
	if err != nil {
		this.Data["json"] = map[string]string{"id": "", "msg": err.Error()}
		x := this.ServeJSON()
		if x != nil {
			log.Println(x)
		}
		return
	}
	this.Data["json"] = "create successfully"
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /:id [put]
func (this *UserController) Put() {
	id, err := this.GetInt(":id")
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	var c models.User
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	log.Println(c)
	if id != c.Id {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.UpdateUser(&c)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = "success"
	}
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /:id [delete]
func (this *UserController) Delete() {
	id, err := this.GetInt(":id")
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.DelUser(&models.User{Id: id})
	if err == nil {
		this.Data["json"] = "delete success!"
	} else {
		this.Data["json"] = "delete failed!"
		this.Ctx.Output.SetStatus(500)
	}
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

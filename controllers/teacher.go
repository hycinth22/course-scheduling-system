package controllers

import (
	"encoding/json"
	"log"

	"courseScheduling/models"
	beego "github.com/beego/beego/v2/server/web"
)

type TeacherController struct {
	beego.Controller
}

// @Title TeacherGetAll
// @Param	search	     query   string  false    "search key"
// @Param	pageIndex	 query   int    false    ""
// @Param	pageSize	 query   int     false    ""
// @Description
// @Success 200 {array} models.Teacher
// @router / [get]
func (this *TeacherController) List() {
	var query struct {
		Search    string `form:"search"`
		PageIndex int    `form:"pageIndex"`
		PageSize  int    `form:"pageSize"`
	}
	if err := this.ParseForm(&query); err != nil {
		log.Println(err)
	}
	var (
		courses []*models.Teacher
		total   int
	)
	if query.Search == "" {
		courses, total = models.ListTeachers(getOffset(query.PageIndex, query.PageSize), query.PageSize)
	} else {
		courses, total = models.SearchTeachers(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Search)
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

// @router /new [post]
func (this *TeacherController) Create() {
	var c models.Teacher
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.AddTeacher(c)
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
func (this *TeacherController) Put() {
	cid := this.GetString(":id")
	if cid != "" {
		var c models.Teacher
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &c)
		if err != nil {
			log.Println(err)
			this.Ctx.Output.SetStatus(400)
			return
		}
		log.Println(c)
		err = models.UpdateTeacher(&c)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = "success"
		}
	}
	err := this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /:id [delete]
func (this *TeacherController) Delete() {
	id := this.GetString(":id")
	err := models.DelTeacher(&models.Teacher{Id: id})
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

// @router /list [get]
func (c *TeacherController) ListAllInColleges() {
	var err error
	col := c.Ctx.Input.Query("college_id")
	if col == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	var r []*models.Teacher
	r, err = models.AllTeachersInColleges(&models.College{Id: col})
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	c.Data["json"] = r
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
}

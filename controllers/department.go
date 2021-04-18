package controllers

import (
	"log"

	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
)

type DepartmentController struct {
	web.Controller
}

// @Title DepartmentGetAll
// @Description get all Department
// @Param	college_id	query 	string	true		"the college_id you want to query its departments"
// @Success 200 {array} models.Department
// @Failure 400 college_id is empty
// @router / [get]
func (c *DepartmentController) GetAll() {
	var err error
	col := c.Ctx.Input.Query("college_id")
	if col == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	var r []*models.Department
	r, err = models.AllDepartmentsInColleges(&models.College{Id: col})
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

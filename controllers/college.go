package controllers

import (
	"log"

	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
)

type CollegeController struct {
	web.Controller
}

// @Title CollegeGetAll
// @Description Get all Colleges
// @Success 200 {array} models.College
// @router / [get]
func (c *CollegeController) GetAll() {
	var err error
	var r []*models.College
	r, err = models.AllColleges()
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

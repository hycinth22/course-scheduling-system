package controllers

import (
	"log"
	"mime/multipart"

	"courseScheduling/excel"
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

// @router /excel [post]
func (c *CollegeController) ImportFromExcel() {
	f, _, err := c.GetFile("collegeExcel")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	batch := excel.ParseCollegeExcel(f)
	err = models.ImportColleges(batch)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		x := c.ServeJSON()
		if x != nil {
			log.Println(x)
			return
		}
		return
	}
}

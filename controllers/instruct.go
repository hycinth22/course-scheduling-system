package controllers

import (
	"encoding/json"
	"log"
	"mime/multipart"

	"courseScheduling/excel"
	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
)

type InstructController struct {
	web.Controller
}

// @router /list [get]
func (this *InstructController) List() {
	var query struct {
		Semester  string `form:"semesterID"`
		Search    string `form:"search"`
		PageIndex int    `form:"pageIndex"`
		PageSize  int    `form:"pageSize"`
	}
	if err := this.ParseForm(&query); err != nil {
		log.Println(err)
	}
	var (
		courses []*models.Instruct
		total   int
	)
	if query.Search == "" {
		courses, total = models.ListInstructs(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Semester)
	} else {
		courses, total = models.SearchInstructs(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Search, query.Semester)
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

// @router /clazzes [get]
func (this *InstructController) GetClazzes() {
	instruct_id, err := this.GetInt("instruct_id")
	if err != nil {
		log.Println(err)
		return
	}
	r, err := models.AllInstructedClazzesForInstruct(instruct_id)
	if err != nil {
		log.Println(err)
		return
	}
	this.Data["json"] = r
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /instructed_clazzes [put]
func (c *InstructController) Put() {
	id, err := c.GetInt("instruct_id")
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	var list []*models.InstructedClazz
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &list)
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	err = models.ResetInstructedClazzForInstruct(&models.Instruct{InstructId: id}, list)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = "success"
	}
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /excel [post]
func (c *InstructController) ImportFromExcel() {
	f, _, err := c.GetFile("excel")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	batch := excel.ParseInstructExcel(f)
	err = models.ImportInstruct(batch)
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

// @router /instructed_clazzes/excel [post]
func (c *InstructController) ImportdInstructClazzesFromExcel() {
	f, _, err := c.GetFile("excel")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	batch := excel.ParseInstructedClazzExcel(f)
	err = models.ImportInstructedClazz(batch)
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

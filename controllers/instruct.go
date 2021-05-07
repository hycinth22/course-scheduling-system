package controllers

import (
	"log"

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

package controllers

import (
	"log"

	"courseScheduling/models"
	beego "github.com/beego/beego/v2/server/web"
)

type SemesterController struct {
	beego.Controller
}

// @Title SemesterGetAll
// @Description Get all Semesters
// @Success 200 {array} models.Semester
// @router / [get]
func (tC *SemesterController) GetAll() {
	var err error
	var r []*models.Semester
	r, err = models.AllSemester()
	if err != nil {
		log.Println(err)
		tC.Ctx.Output.SetStatus(500)
		return
	}
	tC.Data["json"] = r
	err = tC.ServeJSON()
	if err != nil {
		log.Println(err)
		tC.Ctx.Output.SetStatus(500)
		return
	}
}

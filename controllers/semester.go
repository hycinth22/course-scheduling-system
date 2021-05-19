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
func (this *SemesterController) GetAll() {
	var err error
	notHidePast, err := this.GetBool("notHidePast", false)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	var r []*models.Semester
	r, err = models.AllSemester()
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	if !notHidePast && models.GlobalConfig.HiddenPastSemester {
		var filtered []*models.Semester
		for i := range r {
			if !r[i].Expired() {
				filtered = append(filtered, r[i])
			}
			log.Println(r[i].Expired(), r[i])
		}
		this.Data["json"] = filtered
	} else {
		this.Data["json"] = r
	}
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
}

// @router /hide_past_semester_config [get]
func (this *SemesterController) GetPastSemesterVisibility() {
	models.GlobalConfig.RLock()
	this.Data["json"] = map[string]interface{}{
		"val": models.GlobalConfig.HiddenPastSemester,
	}
	models.GlobalConfig.RUnlock()
	err := this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @router /hide_past_semester_config [put]
func (this *SemesterController) ChangePastSemesterVisibility() {
	val, err := this.GetBool("val")
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	models.GlobalConfig.Lock()
	models.GlobalConfig.HiddenPastSemester = val
	models.GlobalConfig.Unlock()
	err = models.GlobalConfig.SaveToDB()
	if err != nil {
		log.Println(err)
	}
}

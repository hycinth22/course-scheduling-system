package controllers

import (
	"encoding/json"
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

// @router /new [post]
func (this *SemesterController) Create() {
	var c models.Semester
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.AddSemester(c)
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
func (this *SemesterController) Put() {
	id := this.GetString(":id")
	var c models.Semester
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	log.Println(c)
	if id != c.StartDate {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.UpdateSemester(&c)
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

// @router / [delete]
func (this *SemesterController) Delete() {
	id := this.GetString("start_date")
	err := models.DelSemester(&models.Semester{StartDate: id})
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

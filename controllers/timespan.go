package controllers

import (
	"log"

	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
)

type TimespanController struct {
	web.Controller
}

// @Title TimespanGetAll
// @Description Get all Timespan
// @Success 200 {array} models.Timespan
// @router / [get]
func (c *TimespanController) GetAll() {
	var err error
	var r []*models.Timespan
	r, err = models.AllTimespan()
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

package controllers

import (
	"log"
	"mime/multipart"

	"courseScheduling/excel"
	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
)

type ClazzController struct {
	web.Controller
}

// @Title ClazzGetAll
// @Description get all Clazz
// @Param	college_id	query 	string	true		"the college_id you want to query its clazz"
// @Success 200 {array} models.Clazz
// @Failure 400 college_id is empty
// @router / [get]
func (c *ClazzController) GetAll() {
	var err error
	col := c.Ctx.Input.Query("college_id")
	if col == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	var r []*models.Clazz
	r, err = models.AllClazzesInColleges(&models.College{Id: col})
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

// @router /list [get]
func (this *ClazzController) List() {
	var query struct {
		Search    string `form:"search"`
		PageIndex int    `form:"pageIndex"`
		PageSize  int    `form:"pageSize"`
	}
	if err := this.ParseForm(&query); err != nil {
		log.Println(err)
	}
	var (
		courses []*models.Clazz
		total   int
	)
	if query.Search == "" {
		courses, total = models.ListClazzes(getOffset(query.PageIndex, query.PageSize), query.PageSize)
	} else {
		courses, total = models.SearchClazzes(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Search)
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

// @router /excel [post]
func (c *ClazzController) ImportFromExcel() {
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
	batch := excel.ParseClazzExcel(f)
	err = models.ImportClazz(batch)
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

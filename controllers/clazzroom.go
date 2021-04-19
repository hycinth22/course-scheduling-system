package controllers

import (
	"encoding/json"
	"log"
	"mime/multipart"

	"courseScheduling/excel"
	"courseScheduling/models"
	beego "github.com/beego/beego/v2/server/web"
)

type ClazzroomController struct {
	beego.Controller
}

// @Title ClazzroomCreate
// @Description create a new Clazzroom
// @Param	body	body 	models.Clazzroom	true		"the Clazzroom body"
// @Success 200 {object} models.Clazzroom
// @Failure 400 invalid request
// @router / [post]
func (c *ClazzroomController) Create() {
	var obj models.Clazzroom
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	err = models.AddClazzroom(&obj)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		x := c.ServeJSON()
		if x != nil {
			log.Println(x)
		}
		return
	}
	c.Data["json"] = obj
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title ClazzroomGetAll
// @Param	search	     query   string  true    "search key"
// @Param	pageIndex	 query   int    true    ""
// @Param	pageSize	 query   int     true    ""
// @Description Get all Clazzrooms
// @Success 200 {array} models.Clazzroom
// @router / [get]
func (c *ClazzroomController) GetAll() {
	var query struct {
		Search    string `form:"search"`
		PageIndex int    `form:"pageIndex"`
		PageSize  int    `form:"pageSize"`
	}
	if err := c.ParseForm(&query); err != nil {
		log.Println(err)
	}
	var (
		result []*models.Clazzroom
		total  int
	)
	if query.Search == "" {
		result, total = models.ListClazzrooms(getOffset(query.PageIndex, query.PageSize), query.PageSize)
	} else {
		result, total = models.SearchClazzrooms(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Search)
	}
	c.Data["json"] = map[string]interface{}{
		"list":      result,
		"pageTotal": total,
	}
	err := c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title Get
// @Description get Clazzroom by cid
// @Param	id	 path   int  true    "Clazzroom id"
// @Success 200 {object} models.Clazzroom
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ClazzroomController) Get() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(400)
		return
	}
	obj, err := models.GetClazzroom(id)
	if err != nil {
		c.Data["json"] = err.Error()
	}
	c.Data["json"] = obj
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title Update
// @Param	cid		    path 	string	true		"The cid you want to update"
// @Param	body		body 	models.Clazzroom	true		"body for content"
// @Success 200 {object} models.Clazzroom
// @Failure 400 :id is not int
// @router /:id [put]
func (c *ClazzroomController) Put() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(400)
		return
	}
	var obj models.Clazzroom
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	obj.Id = id
	log.Println(c)
	err = models.UpdateClazzroom(&obj)
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

// @Title Delete
// @Param	id path    string  true    "Clazzroom id"
// @Success 200 {string} delete success!
// @router /:id [delete]
func (c *ClazzroomController) Delete() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(400)
		return
	}
	err = models.DelClazzroom(&models.Clazzroom{Id: id})
	if err == nil {
		c.Data["json"] = "delete success!"
	} else {
		c.Data["json"] = "delete failed!"
		c.Ctx.Output.SetStatus(500)
	}
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title ClazzroomImportFromExcel
// @Description import Clazzrooms from excel
// @Param	excel   formData    file   true    "the excel file"
// @Success 200 {string} ""
// @Failure 500 import failed
// @router /excel [post]
func (c *ClazzroomController) ImportFromExcel() {
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
	batch := excel.ParseClazzroomExcel(f)
	err = models.ImportClazzroom(batch)
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

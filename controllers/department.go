package controllers

import (
	"encoding/json"
	"log"
	"mime/multipart"

	"courseScheduling/excel"
	"courseScheduling/models"
	"github.com/beego/beego/v2/server/web"
)

type DepartmentController struct {
	web.Controller
}

// @Title DepartmentCreate
// @Description create a new Department
// @Param	body	body 	models.Department	true		"the Department body"
// @Success 200 {object} models.Department
// @Failure 400 invalid request
// @router / [post]
func (c *DepartmentController) Create() {
	var obj models.Department
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	err = models.AddDepartment(&obj)
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

// @Title DepartmentGetAll
// @Description get all Department
// @Param	college_id	query 	string	false		"the college_id you want to query its departments"
// @Success 200 {array} models.Department
// @Failure 400 college_id is empty
// @router / [get]
func (c *DepartmentController) GetAll() {
	var err error
	col := c.Ctx.Input.Query("college_id")
	var r []*models.Department
	if col == "" {
		var query struct {
			Search    string `form:"search"`
			PageIndex int    `form:"pageIndex"`
			PageSize  int    `form:"pageSize"`
		}
		if err := c.ParseForm(&query); err != nil {
			log.Println(err)
			c.Ctx.Output.SetStatus(400)
			return
		}
		var (
			result []*models.Department
			total  int
		)
		if query.Search == "" {
			result, total = models.ListDepartments(getOffset(query.PageIndex, query.PageSize), query.PageSize)
		} else {
			result, total = models.SearchDepartments(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Search)
		}
		c.Data["json"] = map[string]interface{}{
			"list":      result,
			"pageTotal": total,
		}
	} else {
		r, err = models.AllDepartmentsInColleges(&models.College{Id: col})
		if err != nil {
			log.Println(err)
			c.Ctx.Output.SetStatus(500)
			return
		}
		c.Data["json"] = r
	}
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
}

// @Title Get
// @Description get Department by id
// @Param	id	 path   string  true    "Department id"
// @Success 200 {object} models.Department
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DepartmentController) Get() {
	id := c.GetString(":id")
	if id == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	obj, err := models.GetDepartment(id)
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
// @Param	id		    path 	string	true		"The id you want to update"
// @Param	body		body 	models.Clazzroom	true		"body for content"
// @Success 200 {object} models.Department
// @Failure 400 :id is not int
// @router /:id [put]
func (c *DepartmentController) Put() {
	id := c.GetString(":id")
	if id == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	var obj models.Department
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	obj.DeptId = id
	log.Println(c)
	err = models.UpdateDepartment(&obj)
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
// @Param	id path    string  true    "Department id"
// @Success 200 {string} delete success!
// @router /:id [delete]
func (c *DepartmentController) Delete() {
	id := c.GetString(":id")
	if id == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	err := models.DelDepartment(&models.Department{DeptId: id})
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

// @Title DepartmentImportFromExcel
// @Description import Department from excel
// @Param	excel   formData    file   true    "the excel file"
// @Success 200 {string} ""
// @Failure 500 import failed
// @router /excel [post]
func (c *DepartmentController) ImportFromExcel() {
	f, _, err := c.GetFile("departmentExcel")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	batch := excel.ParseDepartmentExcel(f)
	err = models.ImportDepartments(batch)
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

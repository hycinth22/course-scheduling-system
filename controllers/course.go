package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"

	"courseScheduling/excel"
	"courseScheduling/models"

	beego "github.com/beego/beego/v2/server/web"
)

type CourseController struct {
	beego.Controller
}

// @Title CourseCreate
// @Description create a new Course
// @Param	body	body 	models.Course	true		"the course body"
// @Success 200 {int} models.Course.Id
// @Failure 400 invalid parameters
// @router / [post]
func (tC *CourseController) Create() {
	var c models.Course
	err := json.Unmarshal(tC.Ctx.Input.RequestBody, &c)
	if err != nil {
		log.Println(err)
		tC.Ctx.Output.SetStatus(400)
		return
	}
	cid, err := models.AddCourse(c)
	if err != nil {
		tC.Data["json"] = map[string]string{"id": "", "msg": err.Error()}
		x := tC.ServeJSON()
		if x != nil {
			log.Println(x)
		}
		return
	}
	tC.Data["json"] = map[string]string{"id": cid}
	err = tC.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title CourseImportFromExcel
// @Description import Courses from excel
// @Param	excel   formData    file   true    "the excel file"
// @Success 200 {string} ""
// @Failure 500 import failed
// @router /excel [post]
func (tC *CourseController) ImportFromExcel() {
	f, _, err := tC.GetFile("courseExcel")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	batch := excel.ParseCourseExcel(f)
	err = models.ImportCourse(batch)
	if err != nil {
		tC.Ctx.Output.SetStatus(500)
		x := tC.ServeJSON()
		if x != nil {
			log.Println(x)
			return
		}
		return
	}
}

// @Title CourseGetAll
// @Param	search	     query   string  false    "search key"
// @Param	pageIndex	 query   int    false    ""
// @Param	pageSize	 query   int     false    ""
// @Description Get all Courses
// @Success 200 {array} models.Course
// @router / [get]
func (tC *CourseController) GetAll() {
	var query struct {
		Search    string `form:"search"`
		PageIndex int    `form:"pageIndex"`
		PageSize  int    `form:"pageSize"`
	}
	if err := tC.ParseForm(&query); err != nil {
		log.Println(err)
	}
	var (
		courses []*models.Course
		total   int
	)
	if query.Search == "" {
		courses, total = models.ListCourses(getOffset(query.PageIndex, query.PageSize), query.PageSize)
	} else {
		courses, total = models.SearchCourses(getOffset(query.PageIndex, query.PageSize), query.PageSize, query.Search)
	}
	tC.Data["json"] = map[string]interface{}{
		"list":      courses,
		"pageTotal": total,
	}
	err := tC.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title Get
// @Description get Course by cid
// @Param	cid	 path   string  true    "course id"
// @Success 200 {object} models.Course
// @Failure 403 :cid is empty
// @router /:cid [get]
func (tC *CourseController) Get() {
	cid := tC.GetString(":cid")
	if cid != "" {
		err := models.GetCourse(cid)
		if err != nil {
			tC.Data["json"] = err.Error()
		} else {
			tC.Data["json"] = cid
		}
	}
	err := tC.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title Update
// @Param	cid		path 	string	true		"The cid you want to update"
// @Param	body		body 	models.Course	true		"body for Course content"
// @Success 200 {object} models.Course
// @Failure 400 :cid is not int
// @router /:cid [put]
func (tC *CourseController) Put() {
	cid := tC.GetString(":cid")
	if cid != "" {
		var c models.Course
		err := json.Unmarshal(tC.Ctx.Input.RequestBody, &c)
		if err != nil {
			log.Println(err)
			tC.Ctx.Output.SetStatus(400)
			return
		}
		fmt.Println(c)
		err = models.UpdateCourse(&c)
		if err != nil {
			tC.Data["json"] = err.Error()
		} else {
			tC.Data["json"] = "success"
		}
	}
	err := tC.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title Delete
// @Param	cid path    string  true    "the deleted course id"
// @Success 200 {string} delete success!
// @router /:cid [delete]
func (tC *CourseController) Delete() {
	cid := tC.GetString(":cid")
	err := models.DelCourse(cid)
	if err == nil {
		tC.Data["json"] = "delete success!"
	} else {
		tC.Data["json"] = "delete failed!"
		tC.Ctx.Output.SetStatus(500)
	}
	err = tC.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

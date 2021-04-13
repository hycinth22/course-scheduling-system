package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"

	"courseSchduling/excel"
	"courseSchduling/models"

	beego "github.com/beego/beego/v2/server/web"
)

type CourseController struct {
	beego.Controller
}

// @Title CreatCourse
// @Description
// @Param	body		body 	models.Course	true		"body for Course content"
// @Success 200 {int} models.Course.Id
// @Failure 400 body is empty
// @router / [post]
func (tC *CourseController) Post() {
	var c models.Course
	err := json.Unmarshal(tC.Ctx.Input.RequestBody, &tC)
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

// @Title ImportFromExcel
// @Description get all Course
// @Success 200 {int} models.Course.Id
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

// @Title GetAll
// @Description get all Courses
// @Success 200 {object} models.Course
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
// @Param	uid		path 	string	true		"The key for staticblock"
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
// @Description update the Course
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
// @Description delete the Course
// @Param	cid		path 	string	true		"The cid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 cid is empty
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

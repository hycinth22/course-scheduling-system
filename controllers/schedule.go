package controllers

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"courseScheduling/excel"
	"courseScheduling/models"
	"courseScheduling/scheduling"
	"courseScheduling/views"

	beego "github.com/beego/beego/v2/server/web"
)

type ScheduleController struct {
	beego.Controller
}

// @Title ScheduleGetAll
// @Description get all Schedule
// @Param	semester_date	query 	string	true		"the semester you want to query its schedule"
// @Success 200 {array} models.Schedule
// @Failure 400 :semester_date is empty
// @router / [get]
func (c *ScheduleController) GetSchedule() {
	s := c.Ctx.Input.Query("semester_date")
	if s == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	result, err := models.GetSchedulesInSemester(s)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	c.Data["json"] = result
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title ScheduleNew
// @Description create a new Schedule
// @Param	semester_date	query 	string	true		"the semester you want to create its schedule"
// @Success 200 {object} models.Schedule
// @Failure 400 :semester_date is empty
// @router /new [get]
func (c *ScheduleController) NewSchedule() {
	semesterDate := c.Ctx.Input.Query("semester_date")
	if semesterDate == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	semester := &models.Semester{StartDate: semesterDate}
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling(semester)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		log.Println(err)
		return
	}
	allClazzroom, err := models.AllClazzroom()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		log.Println(err)
		return
	}
	allTimespan, err := models.AllTimespan()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		log.Println(err)
		return
	}
	result := scheduling.GenerateSchedule(&scheduling.Params{
		AllInstructedClazz: allInstructedClazz,
		AllClazzroom:       allClazzroom,
		AllTimespan:        allTimespan,
	})
	s, err := models.AddNewSchedule(semester, result, len(allTimespan))
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	c.Data["json"] = s
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
}

// @Title DeleteSchedule
// @Description Delete a Schedule
// @Param	schedule_id	path 	int	true		"the schedule_id you want to delete"
// @Success 200 {string} "ok"
// @Failure 400 :schedule_id is empty
// @router /:schedule_id [delete]
func (c *ScheduleController) DeleteSchedule() {
	id, err := c.GetInt(":schedule_id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		return
	}
	err = models.DelSchedule(id)
	if err != nil {
		c.Data["json"] = err.Error()
	}
	c.Data["json"] = "delete ok!"
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title DeleteAllScheduleInSemester
// @Description Delete All Schedule In Semester
// @Param	semester_date	query 	string	true		"the semester_date you want to delete its's schedules"
// @Success 200 {string} "ok"
// @Failure 400 semester_date is empty
// @router / [delete]
func (c *ScheduleController) DeleteAllScheduleInSemester() {
	semesterDate := c.Ctx.Input.Query("semester_date")
	if semesterDate == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	schs, err := models.GetSchedulesInSemester(semesterDate)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	var ids []int
	for _, s := range schs {
		ids = append(ids, s.Id)
	}
	err = models.DelSchedules(ids)
	if err != nil {
		c.Data["json"] = err.Error()
	}
	c.Data["json"] = "delete ok!"
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title ScheduleGetItems
// @Description get items of Schedule
// @Param	schedule_id	path 	int	true		"the schedule_id you want to query its items"
// @Success 200 {array} models.ScheduleItem
// @Failure 400 :schedule_id is empty
// @router /:schedule_id/items [get]
func (c *ScheduleController) GetScheduleItems() {
	id, err := c.GetInt(":schedule_id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		return
	}
	result, err := models.GetScheduleItems(id)
	if err != nil {
		c.Data["json"] = err.Error()
	}
	c.Data["json"] = result
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

func getScheduleGroupView(id int) (*views.ScheduleItemsTableView, error) {
	sch, err := models.GetSchedule(id)
	if err != nil {
		return nil, err
	}
	if sch == nil {
		return nil, errors.New("not found")
	}
	items, err := models.GetScheduleItems(id)
	if err != nil {
		return nil, err
	}
	result := views.NewScheduleItemsTableView(len(items))
	const cntWeekday = 7
	for _, item := range items {
		// create map
		if result.ByClazz[item.Clazz.ClazzId] == nil {
			result.ByClazz[item.Clazz.ClazzId] = make([]map[int]*models.ScheduleItem, sch.UseTimespan)
			for timespan := 0; timespan < sch.UseTimespan; timespan++ {
				result.ByClazz[item.Clazz.ClazzId][timespan] = make(map[int]*models.ScheduleItem, cntWeekday)
			}
		}
		if result.ByDept[item.Instruct.Teacher.Dept.DeptId] == nil {
			result.ByDept[item.Instruct.Teacher.Dept.DeptId] = make([]map[int][]*models.ScheduleItem, sch.UseTimespan)
			for timespan := 0; timespan < sch.UseTimespan; timespan++ {
				result.ByDept[item.Instruct.Teacher.Dept.DeptId][timespan] = make(map[int][]*models.ScheduleItem, cntWeekday)
			}
		}
		// insert view data
		result.ByClazz[item.Clazz.ClazzId][item.TimespanId-1][item.DayOfWeek] = item
		result.ByDept[item.Instruct.Teacher.Dept.DeptId][item.TimespanId-1][item.DayOfWeek] = append(result.ByDept[item.Instruct.Teacher.Dept.DeptId][item.TimespanId-1][item.DayOfWeek], item)
	}
	return result, nil
}

// @Title ScheduleGetItems
// @Description get items of Schedule
// @Param	schedule_id	path 	int	true		"the schedule_id you want to query its items"
// @Success 200 {object} views.ScheduleItemsTableView
// @Failure 400 :schedule_id is empty
// @router /:schedule_id/items/group_view [get]
func (c *ScheduleController) GetScheduleItemsGroupView() {
	id, err := c.GetInt(":schedule_id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		return
	}
	result, err := getScheduleGroupView(id)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}
	c.Data["json"] = result
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

// @Title ScheduleDownloadExcel
// @Description Schedule DownloadExcel
// @Param	schedule_id	path 	int	true		"the schedule_id you want to download"
// @Param	college_id	query 	string	true		"the college_id you want to download"
// @Success 200 {object} views.ScheduleItemsTableView
// @Failure 400 :schedule_id is empty
// @router /:schedule_id/student_excel [get]
func (c *ScheduleController) ScheduleDownloadStudentExcel() {
	id, err := c.GetInt(":schedule_id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		return
	}
	col := c.Ctx.Input.Query("college_id")
	if col == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	result, err := getScheduleGroupView(id)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}
	var clazzes []*models.Clazz
	clazzes, err = models.AllClazzesInColleges(&models.College{Id: col})
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	tmpfile, err := ioutil.TempFile("", "course_schedule_excel_*.xls")
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}
	defer func(name string) {
		err := tmpfile.Close()
		if err != nil {
			log.Println(err)
			return
		}
		err = os.Remove(name)
		if err != nil {
			log.Println(err)
		}
	}(tmpfile.Name())
	err = excel.GenStudentTables(tmpfile, result, clazzes)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}
	c.Ctx.Output.Download(tmpfile.Name())
}

// @Title ScheduleDownloadExcel2
// @Description Schedule DownloadExcel
// @Param	schedule_id	path 	int	true		"the schedule_id you want to download"
// @Param	college_id	query 	string	true		"the college_id you want to download"
// @Success 200 {object} views.ScheduleItemsTableView
// @Failure 400 :schedule_id is empty
// @router /:schedule_id/teacher_excel [get]
func (c *ScheduleController) ScheduleDownloadTeacherExcel() {
	id, err := c.GetInt(":schedule_id")
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		return
	}
	col := c.Ctx.Input.Query("college_id")
	if col == "" {
		c.Ctx.Output.SetStatus(400)
		return
	}
	result, err := getScheduleGroupView(id)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}
	var depts []*models.Department
	depts, err = models.AllDepartmentsInColleges(&models.College{Id: col})
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	tmpfile, err := ioutil.TempFile("", "course_schedule_excel_*.xls")
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}
	defer func(name string) {
		err := tmpfile.Close()
		if err != nil {
			log.Println(err)
			return
		}
		err = os.Remove(name)
		if err != nil {
			log.Println(err)
		}
	}(tmpfile.Name())
	err = excel.GenTeacherTables(tmpfile, result, depts)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		return
	}
	c.Ctx.Output.Download(tmpfile.Name())
}
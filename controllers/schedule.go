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
func (this *ScheduleController) GetSchedule() {
	s := this.Ctx.Input.Query("semester_date")
	if s == "" {
		this.Ctx.Output.SetStatus(400)
		return
	}
	result, err := models.GetSchedulesInSemester(s)
	if err != nil {
		this.Data["json"] = err.Error()
		return
	}
	this.Data["json"] = result
	err = this.ServeJSON()
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
func (this *ScheduleController) NewSchedule() {
	semesterDate := this.Ctx.Input.Query("semester_date")
	if semesterDate == "" {
		this.Ctx.Output.SetStatus(400)
		return
	}
	semester := &models.Semester{StartDate: semesterDate}
	allInstructedClazz, err := models.AllInstructedClazzesForScheduling(semester)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		log.Println(err)
		return
	}
	allClazzroom, err := models.AllClazzroom()
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		log.Println(err)
		return
	}
	allTimespan, err := models.AllTimespan()
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		log.Println(err)
		return
	}
	result, score := scheduling.GenerateSchedule(&scheduling.Params{
		AllInstructedClazz: allInstructedClazz,
		AllClazzroom:       allClazzroom,
		AllTimespan:        allTimespan,
	})
	s, err := models.AddNewSchedule(semester, result, len(allTimespan), score)
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Data["json"] = s
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
}

// @Title DeleteSchedule
// @Description Delete a Schedule
// @Param	schedule_id	path 	int	true		"the schedule_id you want to delete"
// @Success 200 {string} "ok"
// @Failure 400 :schedule_id is empty
// @router /:schedule_id [delete]
func (this *ScheduleController) DeleteSchedule() {
	id, err := this.GetInt(":schedule_id")
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		return
	}
	err = models.DelSchedule(id)
	if err != nil {
		this.Data["json"] = err.Error()
	}
	this.Data["json"] = "delete ok!"
	err = this.ServeJSON()
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
func (this *ScheduleController) DeleteAllScheduleInSemester() {
	semesterDate := this.Ctx.Input.Query("semester_date")
	if semesterDate == "" {
		this.Ctx.Output.SetStatus(400)
		return
	}
	schs, err := models.GetSchedulesInSemester(semesterDate)
	if err != nil {
		this.Data["json"] = err.Error()
		return
	}
	var ids []int
	for _, s := range schs {
		ids = append(ids, s.Id)
	}
	err = models.DelSchedules(ids)
	if err != nil {
		this.Data["json"] = err.Error()
	}
	this.Data["json"] = "delete ok!"
	err = this.ServeJSON()
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
func (this *ScheduleController) GetScheduleItems() {
	id, err := this.GetInt(":schedule_id")
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		return
	}
	result, err := models.GetScheduleItems(id)
	if err != nil {
		this.Data["json"] = err.Error()
	}
	this.Data["json"] = result
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

func getScheduleGroupView(id int) (*views.ScheduleItemsTableView, float64, error) {
	sch, err := models.GetSchedule(id)
	if err != nil {
		return nil, 0.0, err
	}
	if sch == nil {
		return nil, 0.0, errors.New("not found")
	}
	items, err := models.GetScheduleItems(id)
	if err != nil {
		return nil, 0.0, err
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
		if result.ByTeacherPersonal[item.Instruct.Teacher.Id] == nil {
			result.ByTeacherPersonal[item.Instruct.Teacher.Id] = make([]map[int]*models.ScheduleItem, sch.UseTimespan)
			for timespan := 0; timespan < sch.UseTimespan; timespan++ {
				result.ByTeacherPersonal[item.Instruct.Teacher.Id][timespan] = make(map[int]*models.ScheduleItem, cntWeekday)
			}
		}

		// insert view data
		result.ByClazz[item.Clazz.ClazzId][item.TimespanId-1][item.DayOfWeek] = item
		result.ByDept[item.Instruct.Teacher.Dept.DeptId][item.TimespanId-1][item.DayOfWeek] = append(result.ByDept[item.Instruct.Teacher.Dept.DeptId][item.TimespanId-1][item.DayOfWeek], item)
		result.ByTeacherPersonal[item.Instruct.Teacher.Id][item.TimespanId-1][item.DayOfWeek] = item

	}
	return result, sch.Score, nil
}

// @Title ScheduleGetItems
// @Description get items of Schedule
// @Param	schedule_id	path 	int	true		"the schedule_id you want to query its items"
// @Success 200 {object} views.ScheduleItemsTableView
// @Failure 400 :schedule_id is empty
// @router /:schedule_id/items/group_view [get]
func (this *ScheduleController) GetScheduleItemsGroupView() {
	id, err := this.GetInt(":schedule_id")
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		return
	}
	items, score, err := getScheduleGroupView(id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Data["json"] = map[string]interface{}{
		"score": score,
		"items": items,
	}
	err = this.ServeJSON()
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
func (this *ScheduleController) ScheduleDownloadStudentExcel() {
	id, err := this.GetInt(":schedule_id")
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		return
	}
	col := this.Ctx.Input.Query("college_id")
	if col == "" {
		this.Ctx.Output.SetStatus(400)
		return
	}
	result, _, err := getScheduleGroupView(id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		return
	}
	var clazzes []*models.Clazz
	clazzes, err = models.AllClazzesInColleges(&models.College{Id: col})
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	tmpfile, err := ioutil.TempFile("", "teacher_table_*.xls")
	if err != nil {
		this.Ctx.Output.SetStatus(500)
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
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Ctx.Output.Download(tmpfile.Name())
}

// @Title ScheduleDownloadExcel2
// @Description Schedule DownloadExcel
// @Param	schedule_id	path 	int	true		"the schedule_id you want to download"
// @Param	college_id	query 	string	true		"the college_id you want to download"
// @Success 200 {object} views.ScheduleItemsTableView
// @Failure 400 :schedule_id is empty
// @router /:schedule_id/teacher_excel [get]
func (this *ScheduleController) ScheduleDownloadTeacherExcel() {
	id, err := this.GetInt(":schedule_id")
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		return
	}
	col := this.Ctx.Input.Query("college_id")
	if col == "" {
		this.Ctx.Output.SetStatus(400)
		return
	}
	result, _, err := getScheduleGroupView(id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		return
	}
	var depts []*models.Department
	depts, err = models.AllDepartmentsInColleges(&models.College{Id: col})
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	tmpfile, err := ioutil.TempFile("", "dept_table_*.xls")
	if err != nil {
		this.Ctx.Output.SetStatus(500)
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
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Ctx.Output.Download(tmpfile.Name())
}

// @router /:schedule_id/teacher_personal_excel [get]
func (this *ScheduleController) ScheduleDownloadTeacherPersonalExcel() {
	id, err := this.GetInt(":schedule_id")
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		return
	}
	teacher_id := this.Ctx.Input.Query("teacher_id")
	teacher, err := models.GetTeacher(teacher_id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		return
	}
	result, _, err := getScheduleGroupView(id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		return
	}
	tmpfile, err := ioutil.TempFile("", "teacher_table_*.xls")
	if err != nil {
		this.Ctx.Output.SetStatus(500)
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
	err = excel.GenTeacherPersonalTables(tmpfile, result, teacher)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Ctx.Output.Download(tmpfile.Name())
}

// @router /selected [put]
func (this *ScheduleController) SelectSchedule() {
	semester := this.GetString("semester")
	if semester == "" {
		this.Ctx.Output.SetStatus(400)
		return
	}
	schedule, err := this.GetInt("schedule")
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(400)
		return
	}
	models.GlobalConfig.Lock()
	models.GlobalConfig.SelectedSchedule = schedule
	models.GlobalConfig.SelectedSemester = semester
	models.GlobalConfig.Unlock()
	err = models.GlobalConfig.SaveToDB()
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
	this.Data["json"] = "ok"
	err = this.ServeJSON()
	if err != nil {
		log.Println(err)
		this.Ctx.Output.SetStatus(500)
		return
	}
}

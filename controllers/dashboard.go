package controllers

import (
	"log"
	"math"
	"time"

	"courseScheduling/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/montanaflynn/stats"
)

type DashboardController struct {
	web.Controller
}

type clazzroomuse struct {
	clazzroom int
	timespan  int
	day       int
}

// @router /summary [get]
func (c *DashboardController) GetSummary() {
	models.GlobalConfig.RLock()
	scheduleID := models.GlobalConfig.Config.SelectedSchedule
	semesterDate := models.GlobalConfig.Config.SelectedSemester
	models.GlobalConfig.RUnlock()
	semester, err := models.GetSemester(semesterDate)
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	begin, err := time.Parse("2006/1/2", semester.StartDate)
	weekno := int(math.Ceil(float64(time.Now().Sub(begin)) / float64(7*24*time.Hour)))
	// counts
	cntInstructs, err := models.CountInstructs()
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	cntTeachers, err := models.CountTeachers()
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	cntClazzes, err := models.CountClazzes()
	if err != nil {
		log.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	var (
		tavgLessons, tavgLessonsP float64
		cavgLessons, cavgLessonsP float64
		czusage                   float64
	)
	if scheduleID != 0 {
		sch, err := models.GetSchedule(scheduleID)
		if err == orm.ErrNoRows {
			goto out
		}
		if err != nil {
			return
		}
		view, _, err := getScheduleGroupView(scheduleID)
		if err != nil {
			log.Println(err)
			c.Ctx.Output.SetStatus(500)
			return
		}
		var tlessons []float64
		for _, teacherLessons := range view.ByTeacherPersonal {
			cnt := 0
			for _, r := range teacherLessons {
				for _, t := range r {
					if t != nil {
						cnt++
					}
				}
			}
			tlessons = append(tlessons, float64(cnt))
		}
		tavgLessons, err = stats.Mean(tlessons)
		if err != nil {
			log.Println(err)
			c.Ctx.Output.SetStatus(500)
			return
		}
		var clessons []float64
		for _, clazzLessons := range view.ByClazz {
			cnt := 0
			for _, r := range clazzLessons {
				for _, t := range r {
					if t != nil {
						cnt++
					}
				}
			}
			clessons = append(clessons, float64(cnt))
		}
		cavgLessons, err = stats.Mean(clessons)
		if err != nil {
			log.Println(err)
			c.Ctx.Output.SetStatus(500)
			return
		}
		allclazzroom, err := models.AllClazzroom()
		if err != nil {
			log.Println(err)
			c.Ctx.Output.SetStatus(500)
			return
		}
		clazzroomflag := make(map[clazzroomuse]bool, len(allclazzroom))
		for _, clazzLessons := range view.ByClazz {
			for _, r := range clazzLessons {
				for _, t := range r {
					if t != nil {
						clazzroomflag[clazzroomuse{
							clazzroom: t.Clazzroom.Id,
							timespan:  t.TimespanId,
						}] = true
					}
				}
			}
		}
		cntUnusedClazzroom := 0
		totalCTD := 0
		for _, c := range allclazzroom {
			for timespan := 1; timespan <= sch.UseTimespan; timespan++ {
				for dayofweek := 1; dayofweek <= 5; dayofweek++ {
					totalCTD++
					if !clazzroomflag[clazzroomuse{
						clazzroom: c.Id,
						timespan:  timespan,
					}] {
						cntUnusedClazzroom++
					}
				}
			}
		}
		czusage = (1.0 - float64(cntUnusedClazzroom)/float64(totalCTD)) * 100
		tavgLessonsP = tavgLessons / float64(5*sch.UseTimespan) * 100
		cavgLessonsP = cavgLessons / float64(5*sch.UseTimespan) * 100
	}
out:
	c.Data["json"] = map[string]interface{}{
		"semesterInfo": map[string]interface{}{
			"progress": map[string]interface{}{
				"cur": weekno, "total": semester.Weeks,
			},
			"clazzroom_usage": czusage,
			"teachersStats": map[string]interface{}{
				"percentage": tavgLessonsP,
				"avgLessons": tavgLessons,
			},
			"studentStats": map[string]interface{}{
				"percentage": cavgLessonsP,
				"avgLessons": cavgLessons,
			},
		},
		"counts": map[string]interface{}{
			"lessons":  cntInstructs,
			"teachers": cntTeachers,
			"clazzes":  cntClazzes,
		},
	}
	err = c.ServeJSON()
	if err != nil {
		log.Println(err)
		return
	}
}

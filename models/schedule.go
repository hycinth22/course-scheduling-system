package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Schedule struct {
	// Primary key
	Id          int       `orm:"column(schedule_id);pk;auto" json:"schedule_id"`
	Semester    *Semester `orm:"column(semester_id);rel(fk);type(date)" json:"semester"`
	UseTimespan int       `orm:"column(use_timespan)" json:"use_timespan"`
	UseWeekday  int       `orm:"column(use_weekday)" json:"use_weekday"`
	Score       float64   `orm:"column(score)" json:"score"`
	// Attributes
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

type ScheduleItem struct {
	// Primary key
	ScheduleItemId int `orm:"column(schedule_item_id);pk;auto" json:"-"`
	// Foreign
	ScheduleId int        `orm:"column(schedule_id);index" json:"schedule_id"`
	Instruct   *Instruct  `orm:"column(instruct_id);rel(fk);index" json:"instruct"`
	Clazz      *Clazz     `orm:"column(clazz_id);rel(fk)" json:"clazz"`
	Clazzroom  *Clazzroom `orm:"column(clazzroom_id);rel(fk)" json:"clazzroom"`
	TimespanId int        `orm:"column(timespan_id);" json:"timespan_id"`
	// Attributes
	DayOfWeek int       `orm:"column(day_of_week);" json:"day_of_week"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
}

func (s ScheduleItem) String() string {
	return fmt.Sprintf("课%v 班级%v 教室%v 老师%v", s.Instruct.InstructId, s.Clazz.ClazzId, s.Clazzroom.Id, s.Instruct.Teacher.Id)
}

func GetSchedule(sid int) (*Schedule, error) {
	s := &Schedule{Id: sid}
	err := o.Read(s)
	if err != nil {
		log.Printf("GetSchedule Err: %v\n", err)
	}
	return s, err
}

func DelSchedule(sid int) (err error) {
	s := &Schedule{Id: sid}
	o, err := orm.NewOrm().Begin()
	defer func() {
		if err != nil {
			x := o.Rollback()
			if x != nil {
				log.Println(x)
				return
			}
		}
	}()
	if err != nil {
		log.Printf("DelSchedule: %v\n", err)
	}
	n, err := o.Delete(s)
	if err != nil {
		log.Printf("DelSchedule: %v\n", err)
	}
	n2, err := o.Delete(&ScheduleItem{ScheduleId: sid}, "schedule_id")
	if err != nil {
		return err
	}
	log.Printf("DelSchedule Effected Rows Num: %d %d\n", n, n2)
	err = o.Commit()
	if err != nil {
		log.Printf("DelSchedule: %v\n", err)
	}
	return err
}

func DelSchedules(ids []int) error {
	o, err := orm.NewOrm().Begin()
	if err != nil {
		log.Printf("DelSchedules: %v\n", err)
	}
	defer func() {
		if err != nil {
			x := o.Rollback()
			if x != nil {
				log.Println(x)
				return
			}
		}
	}()
	n, err := o.QueryTable("schedule").Filter("schedule_id__in", ids).Delete()
	if err != nil {
		log.Printf("DelSchedules: %v\n", err)
		return err
	}
	n2, err := o.QueryTable("schedule_item").Filter("schedule_id__in", ids).Delete()
	if err != nil {
		log.Printf("DelSchedules: %v\n", err)
		return err
	}
	log.Printf("DelSchedules Effected Rows Num: %d %d\n", n, n2)
	err = o.Commit()
	if err != nil {
		log.Printf("DelSchedules: %v\n", err)
	}
	return err
}

func AddNewSchedule(semester *Semester, items []*ScheduleItem, useTimespan int, score float64) (s *Schedule, err error) {
	o, err := orm.NewOrm().Begin()
	defer func() {
		if x := recover(); x != nil {
			log.Printf("AddNewSchedule occurred an error:%v. Transcation will rollback\n", err)
			re := o.Rollback()
			if re != nil {
				log.Printf("AddNewSchedule Rollback Error%v\n", re)
			}
			s, err = nil, errors.New(fmt.Sprintln(x))
		}
	}()
	s = new(Schedule)
	s.Semester = semester
	s.UseTimespan = useTimespan
	s.Score = score
	id, err := o.Insert(s)
	if err != nil {
		panic(err)
	}
	for _, item := range items {
		if item.ScheduleId == -1 {
			continue
		}
		item.ScheduleItemId = 0
		item.ScheduleId = int(id)
		_, err := o.Insert(item)
		if err != nil {
			panic(err)
		}
	}
	err = o.Commit()
	if err != nil {
		panic(err)
	}
	return s, nil
}

func GetSchedulesInSemester(semesterDate string) (s []*Schedule, err error) {
	_, err = o.QueryTable("schedule").Filter("semester_id", semesterDate).All(&s)
	return
}

func GetScheduleItems(scheduleID int) (r []*ScheduleItem, err error) {

	num, err := o.QueryTable("schedule_item").Filter("schedule_id", scheduleID).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

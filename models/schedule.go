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
	Id       int       `orm:"column(schedule_id);pk;auto" json:"schedule_id"`
	Semester *Semester `orm:"column(semester_id);rel(fk);type(date)" json:"semester"`
	// Attributes
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type ScheduleItem struct {
	// Primary key
	ScheduleItemId int `orm:"column(schedule_item_id);pk;auto" json:"-"` // Format: "ScheduleId_subId"
	// Foreign
	ScheduleId  int       `orm:"column(schedule_id);index" json:"schedule_id"`
	Instruct    *Instruct `orm:"column(instruct_id);rel(fk);index" json:"instruct_id"`
	Clazz       *Clazz    `orm:"column(clazz_id);rel(fk)" json:"clazz_id"`
	ClazzroomId int       `orm:"column(clazzroom_id);" json:"clazzroom_id"`
	TimespanId  int       `orm:"column(timespan_id);" json:"semester_id"`
	// Attributes
	DayOfWeek int       `orm:"column(day_of_week);" json:"day_of_week"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func (s ScheduleItem) String() string {
	return fmt.Sprintf("课%v 班级%v 教室%v 老师%v", s.Instruct.InstructId, s.Clazz.ClazzId, s.ClazzroomId, s.Instruct.Teacher.Id)
}

func AddNewSchedule(semester *Semester, items []*ScheduleItem) (s *Schedule, err error) {
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
	id, err := o.Insert(s)
	if err != nil {
		panic(err)
	}
	for _, item := range items {
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

package models

import (
	"fmt"
	"time"
)

type Schedule struct {
	// Primary key
	Id int `orm:"column(schedule_id);pk;auto" json:"schedule_id"`
	// Attributes
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type ScheduleItem struct {
	// Primary key
	ScheduleItemId int `orm:"column(schedule_item_id);pk;auto" json:"-"` // Format: "ScheduleId_subId"
	// Foreign
	ScheduleId  *Schedule `orm:"column(schedule_id);rel(fk)" json:"schedule_id"`
	Instruct    *Instruct `orm:"column(instruct_id);rel(fk)" json:"instruct_id"`
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

package models

import "time"

type Schedule struct {
	// Primary key
	Id int `orm:"column(schedule_id);pk;auto" json:"schedule_id"`
	// Attributes
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type ScheduleItem struct {
	// Primary key
	ScheduleId     int `orm:"column(schedule_id)" json:"schedule_id"`
	ScheduleItemId int `orm:"column(schedule_item_id);pk;auto" json:"-"`
	// Foreign keys
	Instruct    *Instruct `orm:"column(instruct_id);rel(fk)" json:"instruct_id"`
	Clazz       *Clazz    `orm:"column(clazz_id);rel(fk)" json:"clazz_id"`
	ClazzroomId string    `orm:"column(clazzroom_id);" json:"clazzroom_id"`
	TimespanId  int       `orm:"column(timespan_id);" json:"semester_id"`
	// Attributes
	DayOfWeek int       `orm:"column(day_of_week);" json:"day_of_week"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

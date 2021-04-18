package models

import "time"

type Student struct {
	// Primary key
	Id string `orm:"column(student_id);pk" json:"student_id"`
	// Attributes
	Name      string    `orm:"column(student_name);" json:"student_name"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
	// Foreign
	Clazz *Clazz `orm:"column(clazz_id);rel(fk)" json:"clazz"`
}

package models

import (
	"time"
)

type Major struct {
	// Primary key
	MajorId string `orm:"column(major_id);pk" json:"major_id"`
	// Attributes
	MajorName string    `orm:"column(major_name)" json:"major_name"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
	// Foreign
	College *College `orm:"column(college_id);rel(fk)" json:"college"`
	// CollegeID string  `orm:"column(college_id)" json:"college_id"`
}

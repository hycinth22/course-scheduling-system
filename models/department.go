package models

import "time"

type Department struct {
	// Primary key
	DeptId string `orm:"column(dept_id);pk" json:"dept_id"`
	// Attributes
	DeptName  string    `orm:"column(dept_name)" json:"dept_name"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
	// Foreign
	College *College `orm:"column(college_id);rel(fk)" json:"college"`
	//CollegeID string `orm:"column(college_id)" json:"college_id"`
}

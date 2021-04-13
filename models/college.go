package models

type College struct {
	// Primary key
	Id string `orm:"column(college_id);pk" json:"college_id"`
	// Attributes
	Name string `orm:"column(college_name);" json:"college_name"`
}

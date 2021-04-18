package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Department struct {
	// Primary key
	DeptId string `orm:"column(dept_id);pk" json:"dept_id"`
	// Attributes
	DeptName  string    `orm:"column(dept_name)" json:"dept_name"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
	// Foreign
	College *College `orm:"column(college_id);rel(fk)" json:"college"`
}

func AllDepartmentsInColleges(coll *College) (r []*Department, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("department").Filter("college_id", coll.Id).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

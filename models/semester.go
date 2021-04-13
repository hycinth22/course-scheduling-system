package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Semester struct {
	// Primary key
	StartDate string `orm:"column(start_date);type(date);pk" json:"start_date"`
	// Attributes
	Name      string    `orm:"column(semester_name);default('')" json:"semester_name"`
	Weeks     int       `orm:"column(semester_weeks)" json:"semester_weeks"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func AddSemester(c Semester) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddSemester %v\n", err)
		return err
	}
	return nil
}

func DelSemester(c *Semester) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelSemester %v\n", err)
	return err
}

func TruncateSemester() error {
	log.Println("TruncateSemester")
	o := orm.NewOrm()
	_, err := o.Raw("truncate table semester").Exec()
	return err
}

func ImportSemester(batch []*Semester) error {
	err := TruncateSemester()
	if err != nil {
		return err
	}
	for _, r := range batch {
		if err := AddSemester(*r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

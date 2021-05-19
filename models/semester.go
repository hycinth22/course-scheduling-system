package models

import (
	"log"
	"time"
)

type Semester struct {
	// Primary key
	StartDate string `orm:"column(start_date);type(date);pk" json:"start_date"`
	// Attributes
	Name      string    `orm:"column(semester_name);default('');index" json:"semester_name"`
	Weeks     int       `orm:"column(semester_weeks)" json:"semester_weeks"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
}

func AllSemester() ([]*Semester, error) {
	var r []*Semester
	num, err := o.QueryTable("semester").All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func AddSemester(c Semester) error {

	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddSemester %v\n", err)
		return err
	}
	return nil
}

func DelSemester(c *Semester) error {

	_, err := o.Delete(c)
	log.Printf("DelSemester %v\n", err)
	return err
}

func TruncateSemester() error {
	log.Println("TruncateSemester")

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

func (s *Semester) Expired() bool {
	begin, err := time.Parse(`2006/1/2`, s.StartDate)
	if err != nil {
		panic(err)
	}
	end := begin.Add(7 * 24 * time.Hour * time.Duration(s.Weeks))
	return end.Before(time.Now())
}

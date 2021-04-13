package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Instruct struct {
	// Primary key
	InstructId int       `orm:"column(instruct_id);pk;auto" json:"-"`
	Teacher    *Teacher  `orm:"column(teacher_id);rel(fk)" json:"teacher"`
	Course     *Course   `orm:"column(course_id);rel(fk)" json:"course"`
	Semester   *Semester `orm:"column(semester_id);rel(fk);type(date)" json:"semester"`
	//TeacherId  string `orm:"column(teacher_id);" json:"teacher_id"`
	//CourseId   string `orm:"column(course_id);" json:"course_id"`
	//SemesterId string `orm:"column(semester_id);" json:"semester_id"`
	// Attributes
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func AddInstruct(c Instruct) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddInstruct %v\n", err)
		return err
	}
	return nil
}

func DelInstruct(c *Instruct) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelInstruct %v\n", err)
	return err
}

func TruncateInstruct() error {
	log.Println("TruncateInstruct")
	o := orm.NewOrm()
	_, err := o.Raw("truncate table instruct").Exec()
	return err
}

func ImportInstruct(batch []*Instruct) error {
	err := TruncateInstruct()
	if err != nil {
		return err
	}
	for _, r := range batch {
		if err := AddInstruct(*r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

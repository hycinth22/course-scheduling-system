package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Teacher struct {
	// Primary key
	Id string `orm:"column(teacher_id);pk" json:"teacher_id"`
	// Attributes
	Name      string    `orm:"column(teacher_name);" json:"teacher_name"`
	Title     string    `orm:"column(teacher_title);" json:"teacher_title"`
	Tel       string    `orm:"column(teacher_tel);" json:"teacher_tel"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
	// Foreign
	Dept *Department `orm:"column(dept_id);rel(fk)" json:"dept"`
}

func AddTeacher(c Teacher) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddTimespan %v\n", err)
		return err
	}
	return nil
}

func DelTeacher(c *Teacher) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelTeacher %v\n", err)
	return err
}

func TruncateTeacher() error {
	log.Println("TruncateTeacher")
	o := orm.NewOrm()
	_, err := o.Raw("truncate table teacher").Exec()
	return err
}

func ImportTeacher(batch []*Teacher) error {
	err := TruncateTeacher()
	if err != nil {
		return err
	}
	for _, r := range batch {
		if err := AddTeacher(*r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

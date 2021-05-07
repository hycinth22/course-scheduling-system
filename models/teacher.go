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

func UpdateTeacher(c *Teacher) error {
	o := orm.NewOrm()
	_, err := o.Update(c)
	if err != nil {
		log.Printf("UpdateTeacher %v\n", err)
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

func ListTeachers(offset, limit int) ([]*Teacher, int) {
	var r []*Teacher
	o := orm.NewOrm()
	num, err := o.QueryTable("teacher").Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("teacher").Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func SearchTeachers(offset, limit int, search string) ([]*Teacher, int) {
	var r []*Teacher
	o := orm.NewOrm()
	cond1 := orm.NewCondition().And("teacher_id__startswith", search).Or("teacher_id__endswith", search)
	cond2 := orm.NewCondition().And("teacher_name__startswith", search).Or("teacher_name__endswith", search)
	cond3 := orm.NewCondition().And("teacher_title__startswith", search).Or("teacher_title__endswith", search)
	cond4 := orm.NewCondition().And("teacher_tel__startswith", search).Or("teacher_tel__endswith", search)
	cond := cond1.OrCond(cond2).OrCond(cond3).OrCond(cond4)
	num, err := o.QueryTable("teacher").SetCond(cond).Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, 0
	}
	cnt, err := o.QueryTable("teacher").SetCond(cond).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
		return nil, 0
	}
	return r, int(cnt)
}

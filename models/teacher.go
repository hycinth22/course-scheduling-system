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

func GetTeacher(id string) (*Teacher, error) {
	t := &Teacher{Id: id}
	err := o.Read(t)
	if err != nil {
		log.Printf("GetTeacher Err: %d, %v\n", err)
	}
	return t, err
}

func AddTeacher(c Teacher) error {
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddTeacher %v\n", err)
		return err
	}
	return nil
}

func UpdateTeacher(c *Teacher) error {

	_, err := o.Update(c)
	if err != nil {
		log.Printf("UpdateTeacher %v\n", err)
		return err
	}
	return nil
}

func DelTeacher(c *Teacher) error {

	_, err := o.Delete(c)
	log.Printf("DelTeacher %v\n", err)
	return err
}

func TruncateTeacher() error {
	log.Println("TruncateTeacher")

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

var o orm.Ormer

func AllTeachersInColleges(coll *College) (r []*Teacher, err error) {
	num, err := o.QueryTable("teacher").Filter("Dept__college_id", coll.Id).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func CountTeachers() (int64, error) {
	cnt, err := o.QueryTable("teacher").Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return cnt, err
}

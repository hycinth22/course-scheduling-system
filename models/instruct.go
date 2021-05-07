package models

import (
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Instruct struct {
	// Alt-Primary key
	InstructId int `orm:"column(instruct_id);pk;auto" json:"instruct_id"`
	// Candidate key
	Teacher  *Teacher  `orm:"column(teacher_id);rel(fk);index" json:"teacher"`
	Course   *Course   `orm:"column(course_id);rel(fk);index" json:"course"`
	Semester *Semester `orm:"column(semester_id);type(date);rel(fk);index" json:"semester"`
	// Attributes
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
}

func (i Instruct) String() string {
	return fmt.Sprintf("Instruct%d {Teacher:%s(%s) Course:%s Semester:%s(%s)}",
		i.InstructId, i.Teacher.Id, i.Teacher.Name, i.Course, i.Semester.StartDate, i.Semester.Name)
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

func ListInstructs(offset, limit int, semester string) ([]*Instruct, int) {
	var r []*Instruct
	o := orm.NewOrm()
	num, err := o.QueryTable("instruct").Filter("semester_id", semester).Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("instruct").Filter("semester_id", semester).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func SearchInstructs(offset, limit int, search string, semester string) ([]*Instruct, int) {
	var r []*Instruct
	o := orm.NewOrm()
	cond1 := orm.NewCondition().And("teacher_id__startswith", search).Or("teacher_id__endswith", search)
	cond2 := orm.NewCondition().And("teacher_name__startswith", search).Or("teacher_name__endswith", search)
	cond3 := orm.NewCondition().And("teacher_title__startswith", search).Or("teacher_title__endswith", search)
	cond4 := orm.NewCondition().And("teacher_tel__startswith", search).Or("teacher_tel__endswith", search)
	cond := cond1.OrCond(cond2).OrCond(cond3).OrCond(cond4)
	num, err := o.QueryTable("instruct").Filter("semester_id", semester).SetCond(cond).Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, 0
	}
	cnt, err := o.QueryTable("instruct").Filter("semester_id", semester).SetCond(cond).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
		return nil, 0
	}
	return r, int(cnt)
}

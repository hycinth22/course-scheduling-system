package models

import (
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type InstructedClazz struct {
	// Primary key
	Id int `orm:"column(id);pk;auto" json:"id"`
	// Candidate key
	Clazz    *Clazz    `orm:"column(clazz_id);rel(fk)" json:"clazz"`
	Instruct *Instruct `orm:"column(instruct_id);rel(fk)" json:"instruct"`
	// Attributes
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func (i InstructedClazz) String() string {
	return fmt.Sprintf("InstructedClazz{Clazz:%s Instruct:%s}", i.Clazz, i.Instruct)
}

func AllInstructedClazzesForScheduling() ([]*InstructedClazz, error) {
	var r []*InstructedClazz
	o := orm.NewOrm()
	num, err := o.QueryTable("instructed_clazz").All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	for i := range r {
		err := o.Read(r[i].Instruct)
		if err != nil {
			return nil, err
		}
		err = o.Read(r[i].Instruct.Course)
		if err != nil {
			return nil, err
		}
		err = o.Read(r[i].Instruct.Teacher)
		if err != nil {
			return nil, err
		}
		err = o.Read(r[i].Clazz)
		if err != nil {
			return nil, err
		}
	}
	return r, err
}

func AddInstructedClazz(c InstructedClazz) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddInstructedClazz %v\n", err)
		return err
	}
	return nil
}

func DelInstructedClazz(c *Instruct) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelInstructedClazz %v\n", err)
	return err
}

func TruncateInstructedClazz() error {
	log.Println("TruncateInstructedClazz")
	o := orm.NewOrm()
	_, err := o.Raw("truncate table instructed_clazz").Exec()
	return err
}

func ImportInstructedClazz(batch []*InstructedClazz) error {
	err := TruncateInstructedClazz()
	if err != nil {
		return err
	}
	for _, r := range batch {
		if err := AddInstructedClazz(*r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

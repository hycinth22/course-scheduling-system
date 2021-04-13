package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type InstructedClazz struct {
	// Alt-Primary key
	Id int `orm:"column(id);pk;auto" json:"id"`
	// Foreign Primary key
	Clazz    *Clazz    `orm:"column(clazz_id);rel(fk)" json:"clazz"`
	Instruct *Instruct `orm:"column(instruct_id);rel(fk)" json:"instruct"`
	// Attributes
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
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
	err := TruncateInstruct()
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

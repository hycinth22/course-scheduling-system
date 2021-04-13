package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Clazzroom struct {
	// Primary key
	Id int `orm:"column(clazzroom_id);pk" json:"-"`
	// Attributes
	Building  string    `orm:"column(building);" json:"building"`
	Room      string    `orm:"column(room);" json:"room"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func AddClazzroom(c Clazzroom) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddClazzroom %v\n", err)
		return err
	}
	return nil
}

func DelClazzroom(c *Clazzroom) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelClazzroom %v\n", err)
	return err
}

func TruncateClazzroom() error {
	log.Println("TruncateClazzroom")
	o := orm.NewOrm()
	_, err := o.Raw("truncate table clazzroom").Exec()
	return err
}

func ImportClazzroom(batch []*Clazzroom) error {
	err := TruncateClazzroom()
	if err != nil {
		return err
	}
	for _, r := range batch {
		if err := AddClazzroom(*r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

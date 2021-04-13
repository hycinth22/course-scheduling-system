package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Clazz struct {
	// Primary key
	ClazzId string `orm:"column(clazz_id);pk" json:"clazz_id"`
	// Attributes
	ClazzName string    `orm:"column(clazz_name)" json:"clazz_name"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
	// Foreign keys
	Major *Major `orm:"column(major_id);rel(fk)" json:"major"`
	// MajorID string `orm:"column(major_id)" json:"major_id"`
}

func AddClazz(c Clazz) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddClazz %v\n", err)
		return err
	}
	return nil
}

func DelClazz(c *Clazz) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelClazz %v\n", err)
	return err
}

func TruncateClazz() error {
	log.Println("TruncateClazz")
	o := orm.NewOrm()
	_, err := o.Raw("truncate table clazz").Exec()
	return err
}

func ImportClazz(batch []*Clazz) error {
	err := TruncateClazz()
	if err != nil {
		return err
	}
	for _, r := range batch {
		if err := AddClazz(*r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

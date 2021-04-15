package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Timespan struct {
	// Primary key
	Id int `orm:"column(timespan_id);pk;auto" json:"timespan_id"`
	// Attributes
	BeginHour   int           `orm:"column(timespan_begin_h)" json:"timespan_begin_h"`
	BeginMinute int           `orm:"column(timespan_begin_m)" json:"timespan_begin_m"`
	EndHour     int           `orm:"column(timespan_end_h)" json:"timespan_end_h"`
	EndMinute   int           `orm:"column(timespan_end_m)" json:"timespan_end_m"`
	Priority    int           `orm:"column(timespan_priority)" json:"timespan_priority"`
	Length      time.Duration `orm:"column(timespan_length);type(bigint)" json:"timespan_length"`
}

func AllTimespan() ([]*Timespan, error) {
	var r []*Timespan
	o := orm.NewOrm()
	num, err := o.QueryTable("timespan").All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func AddOrUpdateTimespan(c *Timespan) error {
	o := orm.NewOrm()
	_, err := o.InsertOrUpdate(c)
	if err != nil {
		log.Printf("AddOrUpdateTimespan %v\n", err)
		return err
	}
	return nil
}

func AddTimespan(c Timespan) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddTimespan %v\n", err)
		return err
	}
	return nil
}

func DelTimespan(c *Timespan) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelCourse %v\n", err)
	return err
}

func TruncateTimespan() error {
	o := orm.NewOrm()
	log.Println("TruncateTimespan")
	_, err := o.Raw("truncate table timespan").Exec()
	return err
}

func ImportTimespan(batch []*Timespan) error {
	err := TruncateTimespan()
	if err != nil {
		return err
	}
	for _, r := range batch {
		if err := AddTimespan(*r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

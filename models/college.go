package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type College struct {
	// Primary key
	Id string `orm:"column(college_id);pk" json:"college_id"`
	// Attributes
	Name string `orm:"column(college_name);" json:"college_name"`
}

func AddOrUpdateCollege(c *College) error {
	o := orm.NewOrm()
	_, err := o.InsertOrUpdate(c)
	if err != nil {
		log.Printf("AddOrUpdateCollege %v\n", err)
		return err
	}
	return nil
}

func AllColleges() ([]*College, error) {
	var r []*College
	o := orm.NewOrm()
	num, err := o.QueryTable("college").All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func TruncateColleges() error {
	o := orm.NewOrm()
	log.Println("TruncateColleges")
	_, err := o.Raw("truncate table college").Exec()
	return err
}

func ImportColleges(batch []*College) error {
	err := TruncateColleges()
	if err != nil {
		return err
	}
	for _, r := range batch {
		err := AddOrUpdateCollege(r)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

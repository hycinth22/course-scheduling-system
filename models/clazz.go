package models

import (
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Clazz struct {
	// Primary key
	ClazzId string `orm:"column(clazz_id);pk" json:"clazz_id"`
	// Attributes
	ClazzName string    `orm:"column(clazz_name)" json:"clazz_name"`
	Spec      string    `orm:"column(spec)" json:"spec"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
	// Foreign
	College *College `orm:"column(college_id);rel(fk)" json:"college"`
}

func (c Clazz) String() string {
	return fmt.Sprintf("Clazz%s(%s)", c.ClazzId, c.ClazzName)
}

func AddClazz(c Clazz) error {

	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddClazz %v\n", err)
		return err
	}
	return nil
}

func UpdateClazz(c *Clazz) error {
	_, err := o.Update(c)
	if err != nil {
		log.Printf("UpdateClazz %v\n", err)
		return err
	}
	return nil
}

func DelClazz(c *Clazz) error {

	_, err := o.Delete(c)
	log.Printf("DelClazz %v\n", err)
	return err
}

func TruncateClazz() error {
	log.Println("TruncateClazz")

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

func AllClazz() ([]*Clazz, error) {
	var r []*Clazz

	num, err := o.QueryTable("clazz").All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func AllClazzesInColleges(coll *College) (r []*Clazz, err error) {

	num, err := o.QueryTable("clazz").Filter("college_id", coll.Id).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func ListClazzes(offset, limit int) ([]*Clazz, int) {
	var r []*Clazz

	num, err := o.QueryTable("clazz").Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("clazz").Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func SearchClazzes(offset, limit int, search string) (r []*Clazz, total int) {

	cond1 := orm.NewCondition().And("clazz_id__startswith", search).Or("clazz_id__endswith", search)
	cond2 := orm.NewCondition().And("clazz_name__startswith", search).Or("clazz_name__endswith", search)
	cond3 := orm.NewCondition().And("college_id__startswith", search).Or("college_id__endswith", search)
	cond := cond1.OrCond(cond2).OrCond(cond3)
	num, err := o.QueryTable("clazz").SetCond(cond).Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, 0
	}
	cnt, err := o.QueryTable("clazz").SetCond(cond).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
		return nil, 0
	}
	return r, int(cnt)
}

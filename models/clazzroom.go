package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Clazzroom struct {
	// Primary key
	Id int `orm:"column(clazzroom_id);pk;auto" json:"id"`
	// Attributes
	Building  string    `orm:"column(building);" json:"building"`
	Room      string    `orm:"column(room);" json:"room"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
}

func ListClazzrooms(offset, limit int) ([]*Clazzroom, int) {
	var r []*Clazzroom
	o := orm.NewOrm()
	num, err := o.QueryTable("clazzroom").OrderBy("building", "room").Offset(offset).Limit(limit).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("clazzroom").Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func SearchClazzrooms(offset, limit int, search string) ([]*Clazzroom, int) {
	var r []*Clazzroom
	o := orm.NewOrm()
	cond1 := orm.NewCondition().And("building__startswith", search).Or("building__endswith", search)
	cond2 := orm.NewCondition().And("room__startswith", search).Or("room__endswith", search)
	cond := cond1.OrCond(cond2)
	num, err := o.QueryTable("clazzroom").SetCond(cond).OrderBy("building", "room").Offset(offset).Limit(limit).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, 0
	}
	cnt, err := o.QueryTable("clazzroom").SetCond(cond).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
		return nil, 0
	}
	return r, int(cnt)
}

func GetClazzroom(id int) (*Clazzroom, error) {
	c := &Clazzroom{Id: id}
	o := orm.NewOrm()
	err := o.Read(c)
	if err != nil {
		log.Printf("GetCourse Err: %d, %v\n", err)
	}
	return c, err
}

func AllClazzroom() ([]*Clazzroom, error) {
	var r []*Clazzroom
	o := orm.NewOrm()
	num, err := o.QueryTable("clazzroom").All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func AddClazzroom(c *Clazzroom) error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	if err != nil {
		log.Printf("AddClazzroom %v\n", err)
		return err
	}
	return nil
}

func UpdateClazzroom(c *Clazzroom) error {
	o := orm.NewOrm()
	_, err := o.Update(c)
	log.Printf("UpdateClazzroom %v\n", err)
	return err
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
		if err := AddClazzroom(r); err != nil {
			log.Println(err)
		}
	}
	return nil
}

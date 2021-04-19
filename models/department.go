package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Department struct {
	// Primary key
	DeptId string `orm:"column(dept_id);pk" json:"dept_id"`
	// Attributes
	DeptName  string    `orm:"column(dept_name)" json:"dept_name"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
	// Foreign
	College *College `orm:"column(college_id);rel(fk)" json:"college"`
}

func AllDepartmentsInColleges(coll *College) (r []*Department, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("department").Filter("college_id", coll.Id).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	return r, err
}

func ListDepartments(offset, limit int) ([]*Department, int) {
	var r []*Department
	o := orm.NewOrm()
	num, err := o.QueryTable("department").RelatedSel().Offset(offset).Limit(limit).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("department").Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func SearchDepartments(offset, limit int, search string) ([]*Department, int) {
	var r []*Department
	o := orm.NewOrm()
	cond1 := orm.NewCondition().And("dept_id__startswith", search).Or("dept_id__endswith", search)
	cond2 := orm.NewCondition().And("dept_name__startswith", search).Or("dept_name__endswith", search)
	cond := cond1.OrCond(cond2)
	num, err := o.QueryTable("department").SetCond(cond).Offset(offset).Limit(limit).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, 0
	}
	cnt, err := o.QueryTable("department").SetCond(cond).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
		return nil, 0
	}
	return r, int(cnt)
}

func GetDepartment(id string) (*Department, error) {
	c := &Department{DeptId: id}
	o := orm.NewOrm()
	err := o.Read(c)
	if err != nil {
		log.Printf("GetDepartment Err: %d, %v\n", err)
	}
	return c, err
}

func AddDepartment(c *Department) error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	if err != nil {
		log.Printf("AddDepartment %v\n", err)
		return err
	}
	return nil
}

func UpdateDepartment(c *Department) error {
	o := orm.NewOrm()
	_, err := o.Update(c)
	log.Printf("UpdateDepartment %v\n", err)
	return err
}

func DelDepartment(c *Department) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	log.Printf("DelDepartment %v\n", err)
	return err
}

func AddOrUpdateDepartment(c *Department) error {
	o := orm.NewOrm()
	_, err := o.InsertOrUpdate(c)
	if err != nil {
		log.Printf("AddOrUpdateDepartment %v\n", err)
		return err
	}
	return nil
}

func TruncateDepartment() error {
	log.Println("TruncateDepartment")
	o := orm.NewOrm()
	_, err := o.Raw("truncate table department").Exec()
	return err
}

func ImportDepartments(batch []*Department) error {
	err := TruncateDepartment()
	if err != nil {
		return err
	}
	for _, r := range batch {
		err := AddOrUpdateDepartment(r)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

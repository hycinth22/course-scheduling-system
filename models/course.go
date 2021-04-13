package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Course struct {
	// Primary key
	Id string `orm:"column(course_id);pk" json:"id"`
	// Attributes
	Name           string    `orm:"column(course_name)" json:"name"`
	Lessons        int       `orm:"column(course_lessons)" json:"lessons"`
	LessonsPerWeek int       `orm:"column(course_lpw)" json:"lessons_per_week"`
	Kind           string    `orm:"column(course_kind);default('')" json:"kind"`
	ExamMode       string    `orm:"column(course_exam_mode);default('')" json:"exam_mode"`
	Founder        string    `orm:"column(course_founder);default('')" json:"founder"`
	CreatedAt      time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt      time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func ListCourses(offset, limit int) ([]*Course, int) {
	var r []*Course
	o := orm.NewOrm()
	num, err := o.QueryTable("course").Offset(offset).Limit(limit).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("course").Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func SearchCourses(offset, limit int, search string) ([]*Course, int) {
	var r []*Course
	o := orm.NewOrm()
	cond1 := orm.NewCondition().And("course_id__startswith", search).Or("course_id__endswith", search)
	cond2 := orm.NewCondition().And("course_name__startswith", search).Or("course_name__endswith", search)
	cond3 := orm.NewCondition().And("course_founder__startswith", search).Or("course_founder__endswith", search)
	cond := cond1.OrCond(cond2).OrCond(cond3)
	num, err := o.QueryTable("course").SetCond(cond).Offset(offset).Limit(limit).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("course").SetCond(cond).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func GetCourse(cid string) error {
	c := &Course{Id: cid}
	o := orm.NewOrm()
	err := o.Read(c)
	if err != nil {
		log.Printf("GetCourse Err: %d, %v\n", err)
	}
	return err
}

func UpdateCourse(c *Course) error {
	o := orm.NewOrm()
	_, err := o.Update(c)
	if err != nil {
		log.Printf("UpdateCourse %v\n", err)
		return err
	}
	return nil
}

func AddOrUpdateCourse(c *Course) error {
	o := orm.NewOrm()
	_, err := o.InsertOrUpdate(c)
	if err != nil {
		log.Printf("AddOrUpdateCourse %v\n", err)
		return err
	}
	return nil
}

func AddCourse(c Course) (string, error) {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddCourse %v\n", err)
		return "", err
	}
	return c.Id, nil
}
func DelCourse(cid string) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Course{Id: cid})
	log.Printf("DelCourse %v\n", err)
	return err
}

func TruncateCourse() error {
	o := orm.NewOrm()
	log.Println("TruncateCourse")
	_, err := o.Raw("truncate table course").Exec()
	return err
}

func ImportCourse(batch []*Course) error {
	err := TruncateCourse()
	if err != nil {
		return err
	}
	for _, r := range batch {
		err := AddOrUpdateCourse(r)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

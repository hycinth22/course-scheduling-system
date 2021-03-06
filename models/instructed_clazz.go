package models

import (
	"fmt"
	"log"
	"time"
)

type InstructedClazz struct {
	// Primary key
	Id int `orm:"column(id);pk;auto" json:"id"`
	// Candidate key
	Clazz    *Clazz    `orm:"column(clazz_id);rel(fk);index" json:"clazz"`
	Instruct *Instruct `orm:"column(instruct_id);rel(fk);index" json:"instruct"`
	// Attributes
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
}

func (i InstructedClazz) String() string {
	return fmt.Sprintf("InstructedClazz{Clazz:%s Instruct:%s}", i.Clazz, i.Instruct)
}

func AllInstructedClazzesForScheduling(semester *Semester) ([]*InstructedClazz, error) {
	var r []*InstructedClazz

	var instructs []int
	_, err := o.Raw("SELECT instruct_id FROM instruct WHERE semester_id = ?", semester.StartDate).QueryRows(&instructs)
	if err != nil {
		log.Printf("Returned Rows: %v\n", err)
		return nil, err
	}
	if len(instructs) == 0 {
		return []*InstructedClazz{}, nil
	}
	num, err := o.QueryTable("instructed_clazz").Filter("instruct_id__in", instructs).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, err
	}
	for i := range r {
		_, err := o.LoadRelated(r[i], "instruct")
		if err != nil {
			return nil, err
		}
		_, err = o.LoadRelated(r[i].Instruct, "course")
		if err != nil {
			return nil, err
		}
		_, err = o.LoadRelated(r[i].Instruct, "teacher")
		if err != nil {
			return nil, err
		}
		_, err = o.LoadRelated(r[i], "clazz")
		if err != nil {
			return nil, err
		}
	}
	return r, err
}

func AllInstructedClazzesForInstruct(instruct_id int) ([]*InstructedClazz, error) {
	var r []*InstructedClazz

	num, err := o.QueryTable("instructed_clazz").Filter("instruct_id", instruct_id).All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, err
	}
	for i := range r {
		_, err := o.LoadRelated(r[i], "instruct")
		if err != nil {
			return nil, err
		}
		_, err = o.LoadRelated(r[i].Instruct, "course")
		if err != nil {
			return nil, err
		}
		_, err = o.LoadRelated(r[i].Instruct, "teacher")
		if err != nil {
			return nil, err
		}
		_, err = o.LoadRelated(r[i], "clazz")
		if err != nil {
			return nil, err
		}
	}
	return r, err
}

func AddInstructedClazz(c InstructedClazz) error {
	_, err := o.Insert(&c)
	if err != nil {
		log.Printf("AddInstructedClazz %v\n", err)
		return err
	}
	return nil
}

func ResetInstructedClazzForInstruct(i *Instruct, cl []*InstructedClazz) error {
	tx, err := o.Begin()
	defer func() {
		if err != nil {
			x := tx.Rollback()
			if x != nil {
				log.Printf("ResetInstructedClazzForInstruct ROLLBACK FAILED %v\n", x)
				return
			}
		} else {
			x := tx.Commit()
			if x != nil {
				log.Printf("ResetInstructedClazzForInstruct COMMIT FAILED %v\n", x)
				return
			}
		}
	}()
	if err != nil {
		log.Printf("ResetInstructedClazzForInstruct %v\n", err)
		return err
	}
	_, err = tx.Delete(&InstructedClazz{Instruct: i}, "instruct_id")
	if err != nil {
		return err
	}
	_, err = tx.InsertMulti(len(cl), cl)
	if err != nil {
		log.Printf("ResetInstructedClazzForInstruct %v\n", err)
		return err
	}
	return nil
}

func DelInstructedClazz(c *Instruct) error {
	_, err := o.Delete(c)
	log.Printf("DelInstructedClazz %v\n", err)
	return err
}

func TruncateInstructedClazz() error {
	log.Println("TruncateInstructedClazz")

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

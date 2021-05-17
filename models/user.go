package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id                int       `orm:"column(id);pk;auto" json:"id"`
	Username          string    `orm:"column(username);unique;index" json:"name"`
	Password          string    `orm:"column(password)" json:"-"`
	Role              string    `orm:"column(role)" json:"role"`
	Status            int       `orm:"column(status)" json:"status"`
	LastLogin         time.Time `orm:"column(last_login_time)" json:"last_login_time"`
	LastLoc           string    `orm:"column(last_login_loc);default:'从未登陆'" json:"last_login_loc"`
	AssociatedTeacher *Teacher  `orm:"column(associated_teacher);null;rel(fk)" json:"associated_teacher"`
	CreatedAt         time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt         time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
}

func CanLogin(username, password string) (bool, *User) {
	u := &User{Username: username}

	err := o.Read(u, "Username")
	if err != nil {
		log.Println("login error", err)
		return false, nil
	}
	if u.Password != password {
		return false, u
	}
	if u.Status != 0 {
		return false, u
	}
	return true, u
}

func UpdateLogin(u *User, loginTime time.Time, loginLocation string) error {
	u.LastLogin = loginTime
	u.LastLoc = loginLocation

	_, err := o.Update(u, "LastLogin", "LastLoc")
	if err != nil {
		log.Println("login error", err)
		return err
	}
	return err
}

func ListUsers(offset, limit int) ([]*User, int) {
	var r []*User

	num, err := o.QueryTable("user").Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
	}
	cnt, err := o.QueryTable("user").Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
	}
	return r, int(cnt)
}

func SearchUsers(offset, limit int, search string) ([]*User, int) {
	var r []*User

	cond1 := orm.NewCondition().And("id__startswith", search).Or("id__endswith", search)
	cond2 := orm.NewCondition().And("username__startswith", search).Or("username__endswith", search)
	cond3 := orm.NewCondition().And("role__startswith", search).Or("role__endswith", search)
	cond4 := orm.NewCondition().And("last_login_loc__startswith", search).Or("last_login_loc__endswith", search)
	cond := cond1.OrCond(cond2).OrCond(cond3).OrCond(cond4)
	num, err := o.QueryTable("user").SetCond(cond).Offset(offset).Limit(limit).RelatedSel().All(&r)
	if err != nil {
		log.Printf("Returned Rows Num: %d, %v\n", num, err)
		return nil, 0
	}
	cnt, err := o.QueryTable("user").SetCond(cond).Count()
	if err != nil {
		log.Printf("Rows Cnt: %d, %v\n", cnt, err)
		return nil, 0
	}
	return r, int(cnt)
}

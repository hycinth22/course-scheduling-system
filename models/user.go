package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id        int       `orm:"pk;auto" json:"id"`
	Username  string    `orm:"unique;index" json:"name"`
	Password  string    `orm:"" json:"pwd"`
	Role      string    `orm:"" json:"role"`
	Status    int       `orm:"" json:"status"`
	LastLogin time.Time `json:"last_login_time"`
	LastLoc   string    `json:"last_login_loc;default:'未登陆'"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
}

func CanLogin(username, password string) (bool, *User) {
	u := &User{Username: username}
	o := orm.NewOrm()
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
	o := orm.NewOrm()
	_, err := o.Update(u, "LastLogin", "LastLoc")
	if err != nil {
		log.Println("login error", err)
		return err
	}
	return err
}

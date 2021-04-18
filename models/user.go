package models

import "time"

type User struct {
	Id        string     `orm:"pk" json:"id"`
	Username  string     `orm:"not null" json:"name"`
	Password  string     `orm:"not null" json:"pwd"`
	Status    int        `json:"status"`
	LastLogin *time.Time `json:"last_login"`
	CreatedAt time.Time  `orm:"column(created_at);auto_now_add;type(datetime)" json:"-"`
	UpdatedAt time.Time  `orm:"column(updated_at);auto_now;type(datetime)" json:"-"`
}

func Login(username, password string) bool {
	return (username == "admin" && password == "123456")
}

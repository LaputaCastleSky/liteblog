package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Name     string `orm:"unique"`
	Email    string
	Avatar   string
	Password string
	Role     int `orm:"default(1)"` // 0 管理员 1正常用户
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Name"},
	}
}

func (u *User) Save() (int64, error) {
	return orm.NewOrm().Insert(u)
}

func (u *User) Exist() bool {
	uu := *u
	err := orm.NewOrm().Read(&uu, "Name")
	if err == orm.ErrNoRows {
		return false
	}
	return true
}

func AuthUser(name string, password string) (*User, bool) {
	u := User{Name: name, Password: password}
	err := orm.NewOrm().Read(&u, "Name", "Password")
	if err == orm.ErrNoRows {
		return nil, false
	}
	return &u, true
}

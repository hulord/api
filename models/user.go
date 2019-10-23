package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	
}


type User struct {
	Id       int
	Username string
	Password string
	Gender   string
	Age      string
	Address  string
	Email    string
}

func GetUserById(id int) (u User, err error) {
	o := orm.NewOrm()
	var user User
	err = o.QueryTable(user).Filter("Id", id).One(&user)
	return user,err
}
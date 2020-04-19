package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	
}


type User struct {
	Id       int 		`json:"id"`
	Username string		`json:"userName"`
	Password string		`json:"Password"`
	Gender   string		`json:"Gender"`
	Age      string		`json:"Age"`
	Address  string		`json:"Address"`
	Email    string		`json:"Email"`
}

func GetUserById(id int) (u User, err error) {
	o := orm.NewOrm()
	var user User
	err = o.QueryTable(user).Filter("Id", id).One(&user)
	return user,err
}

func GetUserByUsername(username string) (u User, err error) {
	o := orm.NewOrm()
	var user User
	err = o.QueryTable(user).Filter("Username", username).One(&user)
	return user,err
}
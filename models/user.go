package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	
}

type User struct {
	Id       int 		`json:"id"`
	Username string		`json:"username"`
	Password string		`json:"password"`
	Gender   string		`json:"gender"`
	Age      string		`json:"age"`
	Address  string		`json:"address"`
	Email    string		`json:"email"`
	Role     int        `json:"role"`
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
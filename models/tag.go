package models

import "github.com/astaxie/beego/orm"
type Tag struct {
	Id int `json:"id"`
	TagName string `json:"tag_name"`
}

func init() {
	orm.RegisterModelWithPrefix("u_db_",new(Tag))
}


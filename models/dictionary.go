package models

import (
	"github.com/astaxie/beego/orm"
)

type Dictionary struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Value string `json:"value"`
	Describe string `json:"describe"`
	Status string `orm:"default(1)";json:"status"`
}

func init() {
	orm.RegisterModelWithPrefix("u_db_",new(Dictionary))
}
func GetDic(t string) (d []Dictionary,err error) {
	o := orm.NewOrm()
	var dic []Dictionary
	if _,err = o.QueryTable(new(Dictionary)).Filter("type__eq",t).All(&dic);err == nil {
		return dic,err
	}
	return dic,nil
}
package models

import (
	"github.com/astaxie/beego/orm"

)

var (
	Menus map[string]*Menu
)

type Role struct {
	Id int `json:"id"`
	Name string `orm:"size(100)"`
}

type RoleMenus struct {
	Id 		int `json:"id"`
	RoleId  int   `json:"role_id"`
	Menu    *Menu  `orm:"rel(fk)"`
}


type Menu struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	ParentId   int    `json:"parent_id"`
	Path       string `json:"path"`
	CreateTime int    `json:"create_time"`
}

func init(){
	orm.RegisterModel(new(Menu),new(Role),new(RoleMenus))

}

func GetMenuByRole(role int) (m []*RoleMenus,err error) {
	orm.Debug = true
	o := orm.NewOrm()
	var arm []*RoleMenus
	o.QueryTable("role_menus").Filter("role_id__eq", 1).RelatedSel().All(&arm)
	return arm,nil
}


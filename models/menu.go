package models

import (
	"github.com/astaxie/beego/orm"
	// "reflect"
	//"encoding/json"
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
	orm.RegisterModelWithPrefix("u_db_",new(Menu),new(Role),new(RoleMenus))
}

func GetMenuByRole(role int64) (m []map[string]interface{},err error) {
	orm.Debug = true
	o := orm.NewOrm()
	var arm []*RoleMenus
	o.QueryTable("u_db_role_menus").Filter("role_id__eq", role).RelatedSel().All(&arm)
	menuList := []map[string]interface{}{}
	for _, value := range arm {
		Menu2 := map[string]interface{}{
			"Id":value.Menu.Id,
			"Name":value.Menu.Name,
			"Icon":value.Menu.Icon,
			"ParentId":value.Menu.ParentId,
			"Path":value.Menu.Path,
			"CreateTime":value.Menu.CreateTime,
		}
		menuList = append(menuList,Menu2)
	}
	return menuList,nil
}


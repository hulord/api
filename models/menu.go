package models

import (

)



type MapRoleMenu struct {
	Id int `json:"id"`
	MenuId int `json:"menu_id"`
	RoleId int `json:"role_id"`
}
	
	
type Menu struct {
	Id         int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Path        string `json:"path"`
	CreateTime int    `json:"createTime"`
}

func (t *Menu) TableName() string {
	return "dc_menu "
}


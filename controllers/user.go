package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"api/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	cid,error := strconv.Atoi(uid)
	if error == nil{
		ob, error := models.GetUserById(cid)
		if error !=nil {
			u.Data["json"] = error
		}else{
			u.Data["json"] = ob
		}
		
	}
	u.ServeJSON()
}




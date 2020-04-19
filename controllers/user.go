package controllers

import (
	"fmt"
	"strconv"
	"api/models"
	"encoding/json"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title Login
// @Description user login
// @Param	username	password 	string	true	
// @Success 200 {object} models.User
// @Failure 403 username or password is wrong
// @router /Login [post]
func (u *UserController) Login(){
	var login_user models.User
	var a interface {} = nil
	data := u.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &login_user)
	if err != nil {
		fmt.Println(err.Error())
	}
	var user models.User
	user,err = models.GetUserByUsername(login_user.Username)
	if err != nil {
		fmt.Println(err.Error())
	}
	if user.Id!=0 {
		if user.Password == login_user.Password {
			u.ApiJsonReturn(0, "登录成功",user)
		}
	}else{	
		u.ApiJsonReturn(1, "无效的用户名和密码",a)	
	}
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





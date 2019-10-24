package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Account
type AccountController struct {
	beego.Controller
}

type returns struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// @Title Login
// @Description account login
// @Param	username	password 	string	true	
// @Success 200 {object} models.User
// @Failure 403 username or password is wrong
// @router /Login [post]
func (a *AccountController) Login(){
	ret := returns{"ok","success",""}
	a.Data["json"] = ret
	a.ServeJSON()
}


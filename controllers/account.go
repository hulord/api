package controllers

import (
	"encoding/json"
	"fmt"
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

type User struct {
	Username string
	Password string
}

// @Title Login
// @Description account login
// @Param	username	password 	string	true	
// @Success 200 {object} models.User
// @Failure 403 username or password is wrong
// @router /Login [post]
func (this *AccountController) Login(){
	var user User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err.Error())
	}
	this.Data["json"] = user.Username
	this.ServeJSON()
}

// @Title Test
// @Description account Test
// @router /Test [get]
func (this *AccountController) Test(){
	fmt.Println("aaa");
}


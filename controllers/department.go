package controllers

import (
	"strconv"
	"api/models"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

// Operations about Users
type DepartmentController struct {
	BaseController
}

// @Title Login
// @Description user login
// @Param	username	password 	string	true	
// @Success 200 {object} models.User
// @Failure 403 username or password is wrong
// @router /login [post]
func (u *UserController) Login(){
	var login_user models.User
	data := u.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &login_user)
	if err != nil {
		u.ApiJsonReturn(1, "无效的用户名和密码4","")
	}

	var user models.User
	user,err = models.GetUserByUsername(login_user.Username)
	if err != nil {
		u.ApiJsonReturn(1, "无效的用户名和密码3","")
	}
	if user.Id!=0 {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(login_user.Password))
		if err ==nil {
			token,_ := CreateToken(user,6000)
			u.ApiJsonReturn(0, "登录成功",token)
		}else{
			u.ApiJsonReturn(1, "无效的用户名和密码1","")
		}
	}else{	
		u.ApiJsonReturn(1, "无效的用户名和密码2","")	
	}
}


	



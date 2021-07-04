package controllers

import (
	"api/models"
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize"
	"golang.org/x/crypto/bcrypt"
	"path"
	"strconv"
	"strings"
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
// @router /login [post]
func (u *UserController) Login() {
	var login_user models.User
	data := u.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &login_user)
	if err != nil {
		u.ApiJsonReturn(1, "无效的用户名和密码", "")
	}

	var user models.User
	user, err = models.GetUserByUsername(login_user.Username)
	if err != nil {
		u.ApiJsonReturn(1, "无效的用户名和密码", "")
	}
	if user.Id != 0 {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login_user.Password))
		if err == nil {
			token, _ := CreateToken(user, 6000)
			u.ApiJsonReturn(0, "登录成功", token)
		} else {
			u.ApiJsonReturn(1, "无效的用户名和密码", "")
		}
	} else {
		u.ApiJsonReturn(1, "无效的用户名和密码", "")
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
	cid, error := strconv.Atoi(uid)
	if error == nil {
		ob, error := models.GetUserById(cid)
		if error != nil {
			u.Data["json"] = error
		} else {
			u.Data["json"] = ob
		}

	}
	u.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (u *UserController) Delete() {
	idStr := u.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUser(id); err == nil {
		u.ApiJsonReturn(0, "删除成功", "")
	} else {
		u.ApiJsonReturn(1, err.Error(), "")
	}
}

// @Title encodePassword
// @Description encode password
// @Success 200 {object} models.User
// @Failure 403 encode is wrong
// @router /encodePassword [get]
func (u *UserController) EncodePassword() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	u.ApiJsonReturn(0, string(hash), "")

}

// 用户列表 ...
// @Title GetDepartment
// @Description get GetUser
// @Success 200 {object} models.Tag
// @Failure 403 :id is empty
// @router /getall [get]
func (c *UserController) GetAll() {
	var fields []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("showCount"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("currentPage"); err == nil {
		offset = v
	}

	//查询权限操作
	if c.role != 1 {
		query["role_id_eq"] = strconv.FormatInt(c.role, 10)
	}
	l, err := models.GetAllUser(query, fields, []string{"Id"}, []string{"desc"}, offset, limit)
	if err != nil {
		c.ApiJsonReturn(1, err.Error(), "")
	} else {
		c.ApiJsonReturn(0, "成功", l)
	}
}

// 导入Excel ...
// @Title GetDepartment
// @Description get GetDepartment
// @Success 200 {object} models.Tag
// @Failure 403 :id is empty
// @router /import [get]
func (c *UserController) Import() {
	file, h, _ := c.GetFile("file") //获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	AllowExtMap := map[string]bool{
		".xlsx": true,
		".xls":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		c.ApiJsonReturn(1, "后缀名不符合上传要求", "")
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		c.ApiJsonReturn(1, err.Error(), "")
	}
	rows := xlsx.GetRows("Sheet1")
	for i, row := range rows {
		if i > 0 {
			var userStruct models.User
			for _, _ = range row {
				userStruct.Username = row[0]
				userStruct.Password = row[1]
				userStruct.Age = row[2]
				userStruct.Address = row[3]
				roleId, _ := strconv.Atoi(row[4])
				userStruct.Role = &models.Role{Id: roleId}

				tagId, _ := strconv.Atoi(row[5])
				userStruct.Tag = &models.Dictionary{Id: tagId}
			}
			if _, err := models.AddUser(&userStruct); err != nil {
				c.ApiJsonReturn(1, err.Error(), "")
			}
		}
	}
	c.ApiJsonReturn(0, "成功", "")
}

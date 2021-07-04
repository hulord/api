package controllers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"path"
	"reflect"
	"strconv"
	"strings"
)

// Operations about Attendance
type AttendanceController struct {
	BaseController
}

// URLMapping ...
func (c *AttendanceController) URLMapping() {
	c.Mapping("getone", c.GetOne)
	c.Mapping("getall", c.GetAll)
	c.Mapping("review", c.Review)
	c.Mapping("applyList", c.ApplyList)
}

// GetOne ...
// @Title Get One
// @Description get Artical by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Artical
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AttendanceController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLeaveById(id)
	if err != nil {
		c.ApiJsonReturn(1, "内容不存在", "")
	} else {
		c.ApiJsonReturn(0, "", v)
	}
}

// Put ...
// @router /:id [put]
func (c *AttendanceController) Put() {
	var leaveStruct models.Leave
	idStr := c.Ctx.Input.Param(":id")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &leaveStruct); err == nil {
		id, _ := strconv.Atoi(idStr)
		leaveStruct.User = &models.User{Id: int(c.userId)}
		leaveStruct.DealUser = &models.User{Id: int(c.userId)}
		if err := models.UpdateLeaveById(id, &leaveStruct); err == nil {
			c.ApiJsonReturn(0, "更新成功", "")
		} else {
			c.ApiJsonReturn(1, err.Error(), "")
		}
	} else {
		c.ApiJsonReturn(1, err.Error(), "")
	}
}

// Add ...
// @Title Add
// @Description Add Artical by data
// @Param	title	query	string	true	"artical title. e.g. 文章标题"
// @Param	author	query	string		"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Artical
// @Failure 403 :id is empty
// @router /create [Post]
func (c *AttendanceController) Create() {
	var leaveStruct models.Leave
	//验证字符是否为空
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &leaveStruct); err == nil {
		leaveStruct.User = &models.User{Id: int(c.userId)}
		leaveStruct.DealUser = &models.User{Id: int(c.userId)}
		if _, err := models.AddLeave(&leaveStruct); err == nil {
			c.ApiJsonReturn(0, "新建成功", leaveStruct)
		} else {
			c.ApiJsonReturn(1, err.Error(), "")
		}
	} else {
		c.ApiJsonReturn(1, err.Error(), "")
	}
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AttendanceController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)
	if err := models.DeleteAttendance(id); err == nil {
		c.ApiJsonReturn(0, "删除成功", "")
	} else {
		c.ApiJsonReturn(1, err.Error(), "")
	}
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /review/:id [Post]
func (c *AttendanceController) Review() {
	var attendanceStruct models.Attendance
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &attendanceStruct); err == nil {

		c.ApiJsonReturn(1, err.Error(), "")
	}
}

// department ...
// @Title GetDepartment
// @Description get GetDepartment
// @Success 200 {object} models.Tag
// @Failure 403 :id is empty
// @router /getall [get]
func (c *AttendanceController) GetAll() {
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

	//新增查询权限操作
	if c.role == 2 {
		query["user_id"] = strconv.FormatInt(c.userId, 10)
	}
	if c.role == 3 {
		query["user_id"] = strconv.FormatInt(c.userId, 10)
	}
	l, err := models.GetAllAttendance(query, fields, []string{"Id"}, []string{"desc"}, offset, limit)
	if err != nil {
		c.ApiJsonReturn(1, err.Error(), "")
	} else {
		ml := make([]interface{}, len(l.DataList))
		fmt.Println(len(l.DataList))
		if len(l.DataList) > 0 {
			for k, v := range l.DataList {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)

				m["id"] = val.FieldByName("Id").Interface()
				m["attendance"] = utils.JSONToMap(utils.Strval(val.FieldByName("Attendance").Interface()))
				m["user"] = val.FieldByName("User").Interface()
				m["mouth"] = val.FieldByName("Mouth").Interface()
				m["create_time"] = val.FieldByName("CreateTime").Interface()

				ml[k] = m
			}
			l.DataList = ml
		}
		c.ApiJsonReturn(0, "成功", l)
	}
}

// department ...
// @Title GetDepartment
// @Description get GetDepartment
// @Success 200 {object} models.Tag
// @Failure 403 :id is empty
// @router /applyList [get]
func (c *AttendanceController) ApplyList() {
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

	//新增查询权限操作
	if c.role == 2 {
		query["role"] = strconv.FormatInt(c.role, 10)
	}
	if c.role == 3 {
		query["user_id"] = strconv.FormatInt(c.userId, 10)
	}
	l, err := models.GetAllLeave(query, fields, []string{"Id"}, []string{"desc"}, offset, limit)
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
func (c *AttendanceController) Import() {
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
			var attendanceStruct models.Attendance
			for _, _ = range row {
				mouth, _ := strconv.Atoi(row[1])
				attendanceStruct.Mouth = mouth
				attendanceStruct.Attendance = row[2]
				userId, _ := strconv.Atoi(row[0])
				attendanceStruct.User = &models.User{Id: userId}
			}
			if _, err := models.AddAttendance(&attendanceStruct); err != nil {
				c.ApiJsonReturn(1, err.Error(), "")
			}
		}
	}
	c.ApiJsonReturn(0, "成功", "")
}

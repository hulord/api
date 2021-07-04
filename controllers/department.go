package controllers

import (
	"api/models"
	"strings"
	"strconv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"path"
)

// Operations about Department
type DepartmentController struct {
	BaseController
}

// URLMapping ...
func (c *DepartmentController) URLMapping() {
	c.Mapping("getall", c.GetAll)
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DepartmentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println(id)
	if err := models.DeleteDepartment(id); err == nil {
		c.ApiJsonReturn(0, "删除成功", "")
	} else {
		c.ApiJsonReturn(1, err.Error(), "")
	}
}

// department ...
// @Title GetDepartment
// @Description get GetDepartment
// @Success 200 {object} models.Tag
// @Failure 403 :id is empty
// @router /getall [get]
func (c *DepartmentController) GetAll() {
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
	// query: k:v,k:v
	query["type.eq"] = "department"
	query["status.eq"] = "1"

	//新增查询权限操作
	//if  c.role != 0{
	//	query["role_id"] = strconv.FormatInt(c.role,10)
	//}
	l, err := models.GetAllDic(query, fields, []string{"Id"}, []string{"desc"}, offset, limit)
	if err != nil {
		c.ApiJsonReturn(1, err.Error(), "")
	} else {
		c.ApiJsonReturn(0, "成功", l)
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
			var dictionaryStruct models.Dictionary
			for _, _ = range row {

				dictionaryStruct.Name = row[0]
				dictionaryStruct.Type = "department"
				dictionaryStruct.Describe = row[3]
				dictionaryStruct.Status = 1
			}
			if _, err := models.AddDepartment(&dictionaryStruct); err != nil {
				c.ApiJsonReturn(1, err.Error(), "")
			}
		}
	}
	c.ApiJsonReturn(0, "成功", "")
}
}



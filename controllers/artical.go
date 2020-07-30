package controllers

import (
	"api/models"
	"encoding/json"
	"strconv"
	"strings"
)

// ArticalController operations for Artical
type ArticalController struct {
	BaseController
}

// URLMapping ...
func (c *ArticalController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Test1",c.Test1)
}

// Post ...
// @Title Post
// @Description create Artical
// @Param	body		body 	models.Artical	true		"body for Artical content"
// @Success 201 {int} models.Artical
// @Failure 403 body is empty
// @router / [post]
func (c *ArticalController) Post() {
	var v models.Artical
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddArtical(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Artical by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Artical
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ArticalController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetArticalTagsById(id)
	if err != nil {
		c.ApiJsonReturn(1,err.Error(),"")	
	} else {
		c.ApiJsonReturn(0,"",v)	
	}
}


// GetAll ...
// @Title Get All
// @Description get Artical
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Artical
// @Failure 403
// @router /getall [get]
func (c *ArticalController) GetAll() {
	var fields []string
	var sortby []string
	var order  []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("pageSize"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("page"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	 if v := c.GetString("keyword"); v != "" {
		 query["title.contains"] = v
	// 	for _, cond := range strings.Split(v, ",") {
	// 		kv := strings.SplitN(cond, ":", 2)
	// 		if len(kv) != 2 {
	// 			c.ApiJsonReturn(0,"","Error: invalid query key/value pair")	
	// 		}
	// 		k, v := kv[0], kv[1]
	// 		query[k] = v
	// 	}
	}
	
	//新增查询权限操作
	if c.role != 0{
		query["role"] = strconv.Itoa(c.role)
	}

	l, err := models.GetAllArtical(query, fields, sortby, order,offset, limit)
	if err != nil {
		c.ApiJsonReturn(1,err.Error(),"")	
	} else {
		c.ApiJsonReturn(0,"",l)	
	}

}

// Put ...
// @Title Put
// @Description update the Artical
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Artical	true		"body for Artical content"
// @Success 200 {object} models.Artical
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArticalController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Artical{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateArticalById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Artical
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArticalController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteArtical(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Test ...
// @Title Delete
// @Description delete the Artical
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /test/:id [get]
func (c *ArticalController) Test() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	l, err := models.GetArticalTags2ById(id)
	if err != nil {
		c.ApiJsonReturn(1,err.Error(),"")	
	} else {
		c.ApiJsonReturn(0,"",l)	
	}
}

// Test ...
// @Title Delete
// @Description delete the Artical
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /Test1 [Post]
func (c *ArticalController) Test1() {
	c.ApiJsonReturn(0,"",1)	
}
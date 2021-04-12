package controllers

import (
	"api/models"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//  ImageController operations for Image
type ImageController struct {
	BaseController
}

// URLMapping ...
func (c *ImageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Upload", c.UploadFile)
}

// Post ...
// @Title Post
// @Description create Image
// @Param	body		body 	models.Image	true		"body for Image content"
// @Success 201 {int} models.Image
// @Failure 403 body is empty
// @router / [post]
func (c *ImageController) Post() {
	var v models.Image
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddImage(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Image by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Image
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ImageController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetImageById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Image
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Image
// @Failure 403
// @router / [get]
func (c *ImageController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
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
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllImage(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.ApiJsonReturn(0, err.Error(), "")
	} else {
		c.ApiJsonReturn(0, "成功", l)
	}
}

// Put ...
// @Title Put
// @Description update the Image
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Image	true		"body for Image content"
// @Success 200 {object} models.Image
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ImageController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Image{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateImageById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Image
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ImageController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteImage(id); err == nil {
		c.ApiJsonReturn(0, "删除成功", "")
	} else {
		c.ApiJsonReturn(1, "删除失败", "")
	}
}

// upload file ...
// @Title Delete
// @Description delete the Artical
// @Param	File 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /Upload [Post]
func (c *ImageController) UploadFile() {
	f, h, _ := c.GetFile("file") //获取上传的文件
	ext := path.Ext(h.Filename)

	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		c.ApiJsonReturn(1, "上传失败", "")
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext
	//this.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("file", fpath)
	if err != nil {
		c.ApiJsonReturn(1, "上传失败", "")
	}
	//写入图片数据表
	var imageInfo models.Image
	ctype := c.GetString("type")
	if ctype == "" {
		c.ApiJsonReturn(1, "参数错误", "")
	}
	imageInfo.Name = fileName
	imageInfo.Type = ctype
	imageInfo.Url = "http://" + c.Ctx.Request.Host + "/" + fpath
	_, err = models.AddImage(&imageInfo)
	if err != nil {
		c.ApiJsonReturn(1, "上传失败", "")
	}
	c.ApiJsonReturn(0, "上传成功", imageInfo)
}

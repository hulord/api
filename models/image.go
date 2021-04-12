package models

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Image struct {
	Id         int64     `json:"uid"`
	Name       string    `json:"name" orm:"size(128)"`
	Type       string    `orm:"size(50)" json:"-"`
	Url        string    `json:"url" orm:"size(128)"`
	CreateTime time.Time `json:"-"          orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `json:"-"          orm:"auto_now";type(datetime)`
}

type List struct {
	ShowCount   int64         `json:"showCount"`
	CurrentPage int64         `json:"currentPage"`
	TotalResult int64         `json:"totalResult"`
	DataList    []interface{} `json:"dataList"`
}

func init() {
	orm.RegisterModelWithPrefix("u_db_", new(Image))
}

// AddImage insert a new Image into database and returns
// last inserted Id on success.
func AddImage(m *Image) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetImageById retrieves Image by Id. Returns error if
// Id doesn't exist
func GetImageById(id int64) (v *Image, err error) {
	o := orm.NewOrm()
	v = &Image{Id: id}
	if err = o.QueryTable(new(Image)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllImage retrieves all Image matches certain condition. Returns empty list if
// no records exist
func GetAllImage(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (listPage List, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Image))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	qs = qs.RelatedSel()
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return listPage, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return listPage, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return listPage, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return listPage, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Image
	qs = qs.OrderBy(sortFields...)
	var ml []interface{}

	if count, err := qs.Count(); err == nil {
		listPage.TotalResult = count
	} else {
		listPage.TotalResult = 0
	}
	listPage.CurrentPage = offset
	listPage.ShowCount = limit

	offset = (offset - 1) * limit
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		listPage.DataList = ml
		return listPage, nil
	}
	return listPage, err
}

// UpdateImage updates Image by Id and returns error if
// the record to be updated doesn't exist
func UpdateImageById(m *Image) (err error) {
	o := orm.NewOrm()
	v := Image{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteImage deletes Image by Id and returns error if
// the record to be deleted doesn't exist
func DeleteImage(id int64) (err error) {
	o := orm.NewOrm()
	v := Image{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Image{Id: id}); err == nil {
			if err = DeleteFile(v.Url); err == nil {
				fmt.Println("Number of records deleted in database:", num)
			}
		}
	}
	return
}

func DeleteFile(path string) (err error) {
	comma := strings.Index(path, "static")
	path = path[comma:]
	err = os.Remove(path) //删除文件
	if err != nil {
		//删除失败,输出错误详细信息
		return err
	}
	return
}

package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
)
type ArticalList struct{
	PageSize int64 `json:"pagesize"`
	Page int64 `json:"page"`
	Total int64 `json:"total"`
	List []interface{} `json:"list"`
}

type Tag struct {
	Id int `json:"id"`
	TagName string `json:"tag_name"`
}
	
type Artical struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int    `json:"createTime"`
	Role       int64  `json:"role"`
	Tags       []*Tag `orm:"rel(m2m)"`
}

func (t *Artical) TableName() string {
	return "artical"
}

func init() {
	orm.RegisterModel(new(Artical),new(Tag))
}

// AddArtical insert a new Artical into database and returns
// last inserted Id on success.
func AddArtical(m *Artical) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticalById retrieves Artical by Id. Returns error if
// Id doesn't exist
func GetArticalById(id int) (v *Artical, err error) {
	o := orm.NewOrm()
	v = &Artical{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetArticalById retrieves Artical by Id. Returns error if
// Id doesn't exist
func GetArticalTagsById(id int) (v Artical,err error) {
	orm.Debug = true
	o := orm.NewOrm()
	artical := Artical{Id: id}

	if err := o.Read(&artical); err == nil {
		o.LoadRelated(&artical, "Tags")
	}
	return artical,nil		
}


// GetArticalById retrieves Artical by Id. Returns error if
// Id doesn't exist
func GetArticalTags2ById(id int) (v Artical,err error) {
	orm.Debug = true
	o := orm.NewOrm()
	artical := Artical{Id: id}
	if err := o.Read(&artical); err == nil {
		o.LoadRelated(&artical, "Tags")
	}
	return artical,nil		
}


// GetAllArtical retrieves all Artical matches certain condition. Returns empty list if
// no records exist
func GetAllArtical(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (a ArticalList, err error) {
	var articalList ArticalList
	o := orm.NewOrm()
	qs := o.QueryTable(new(Artical))
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
					return articalList, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return articalList, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return articalList, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return articalList, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Artical
	qs = qs.OrderBy(sortFields...)
	var ml []interface{}
	
	if count,err := qs.Count(); err == nil {
		articalList.Total = count
	} else {
		articalList.Total = 0
	}
	articalList.Page = offset
	articalList.PageSize = limit

	offset =  (offset-1)*limit
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
		articalList.List = ml
		return articalList, nil
	}
	return articalList, err
}

// UpdateArtical updates Artical by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticalById(m *Artical) (err error) {
	o := orm.NewOrm()
	v := Artical{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArtical deletes Artical by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArtical(id int) (err error) {
	o := orm.NewOrm()
	v := Artical{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Artical{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

type DictionaryList struct {
	ShowCount   int64         `json:"showCount"`
	CurrentPage int64         `json:"currentPage"`
	TotalResult int64         `json:"totalResult"`
	DataList    []interface{} `json:"dataList"`
}

type Dictionary struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Describe string `json:"describe"`
	Status   string `json:"status";orm:"default(1)"`
}

func init() {
	orm.RegisterModelWithPrefix("u_db_", new(Dictionary))
}

func AddDepartment(m *Dictionary) (id int64, err error) {
	inserter, _ := orm.NewOrm().QueryTable(new(Dictionary)).PrepareInsert()
	o := orm.NewOrm()
	if id, err = o.Insert(m); err == nil {
		inserter.Close()
	}
	return id, err
}

func GetDic(t string) (d []Dictionary, err error) {
	o := orm.NewOrm()
	var dic []Dictionary
	if _, err = o.QueryTable(new(Dictionary)).Filter("type__eq", t).Filter("status", 1).All(&dic); err == nil {
		return dic, err
	}
	return dic, nil
}

func DeleteDepartment(id int) (err error) {
	o := orm.NewOrm()
	v := Dictionary{Id: id}
	fmt.Println(v)
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Delete(&v); err == nil {
			return err
		}
	}
	return err
}

func GetAllDic(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (a DictionaryList, err error) {
	var dictionaryList DictionaryList
	o := orm.NewOrm()
	qs := o.QueryTable(new(Dictionary))
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
					return dictionaryList, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return dictionaryList, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return dictionaryList, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return dictionaryList, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Dictionary
	qs = qs.OrderBy(sortFields...)
	var ml []interface{}

	if count, err := qs.Count(); err == nil {
		dictionaryList.TotalResult = count
	} else {
		dictionaryList.TotalResult = 0
	}
	dictionaryList.CurrentPage = offset
	dictionaryList.ShowCount = limit

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
		dictionaryList.DataList = ml
		return dictionaryList, nil
	}
	return dictionaryList, err
}

func deleteDic(t string) (d []Dictionary, err error) {
	o := orm.NewOrm()
	var dic []Dictionary
	if _, err = o.QueryTable(new(Dictionary)).Filter("type__eq", t).Filter("status", 1).All(&dic); err == nil {
		return dic, err
	}
	return dic, nil
}

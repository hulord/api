package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

func init() {

}

type UserList struct {
	ShowCount   int64         `json:"showCount"`
	CurrentPage int64         `json:"currentPage"`
	TotalResult int64         `json:"totalResult"`
	DataList    []interface{} `json:"dataList"`
}

type User struct {
	Id       int         `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Gender   string      `json:"gender"`
	Age      string      `json:"age"`
	Address  string      `json:"address"`
	Email    string      `json:"email"`
	Salf     string      `json:"salf"`
	Role     *Role       `json:"role" orm:"rel(fk);null,on_delete(set_null)"`
	Tag      *Dictionary `json:"department" orm:"rel(fk);null;on_delete(set_null)"`
}

func AddUser(m *Dictionary) (id int64, err error) {
	inserter, _ := orm.NewOrm().QueryTable(new(User)).PrepareInsert()
	o := orm.NewOrm()
	if id, err = o.Insert(m); err == nil {
		inserter.Close()
	}
	return id, err
}

func GetUserById(id int) (u User, err error) {
	o := orm.NewOrm()
	var user User
	err = o.QueryTable(user).Filter("Id", id).One(&user)
	return user, err
}

func GetUserByUsername(username string) (u User, err error) {
	o := orm.NewOrm()
	var user User

	err = o.QueryTable(user).Filter("Username", username).RelatedSel().One(&user)
	return user, err
}

func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Delete(&v); err == nil {
			return err
		}
	}
	return err
}

func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (a UserList, err error) {
	var userList UserList
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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
					return userList, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return userList, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return userList, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return userList, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...)
	var ml []interface{}

	if count, err := qs.Count(); err == nil {
		userList.TotalResult = count
	} else {
		userList.TotalResult = 0
	}
	userList.CurrentPage = offset
	userList.ShowCount = limit

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
		userList.DataList = ml
		return userList, nil
	}
	return userList, err
}

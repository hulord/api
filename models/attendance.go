package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

type Attendance struct {
	Id         int       `json:"id"`
	User       *User     `json:"user" orm:"rel(fk);null,on_delete(set_null)"`
	Mouth      int       `json:"mouth"`
	Attendance string    `json:"attendance"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

type Leave struct {
	Id         int       `json:"id"`
	User       *User     `json:"user" orm:"rel(fk);null,on_delete(set_null)"`
	Type       int       `json:"type"`
	Reason     string    `json:"reason"`
	DealReason string    `json:"deal_reason"`
	Result     int       `json:"result"`
	StartTime  time.Time `json:"start_time" orm:"type(datetime)"`
	EndTime    time.Time `json:"end_time"   orm:"type(datetime)"`
	DealUser   *User     `json:"deal_user"  orm:"rel(fk);null,on_delete(set_null)"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	Status     int       `json:"status"`
}

func init() {
	orm.RegisterModelWithPrefix("u_db_", new(Attendance), new(Leave))
}

type LeavePageList struct {
	ShowCount   int64         `json:"showCount"`
	CurrentPage int64         `json:"currentPage"`
	TotalResult int64         `json:"totalResult"`
	DataList    []interface{} `json:"dataList"`
}

type AttendancePageList struct {
	ShowCount   int64         `json:"showCount"`
	CurrentPage int64         `json:"currentPage"`
	TotalResult int64         `json:"totalResult"`
	DataList    []interface{} `json:"dataList"`
}

func GetLeaveById(id int) (v *Leave, err error) {
	orm.Debug = true
	o := orm.NewOrm()
	leave := &Leave{}
	o.QueryTable(new(Leave)).Filter("Id", id).RelatedSel().One(leave)
	if err == nil {
		return leave, nil
	}
	return nil, err
}

func AddAttendance(m *Attendance) (id int64, err error) {
	inserter, _ := orm.NewOrm().QueryTable(new(Attendance)).PrepareInsert()
	o := orm.NewOrm()
	if id, err = o.Insert(m); err == nil {
		inserter.Close()
	}
	return id, err
}

// UpdateArtical updates Artical by Id and returns error if
// the record to be updated doesn't exist
func UpdateLeaveById(id int, m *Leave) (err error) {
	o := orm.NewOrm()
	v := Leave{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return err
}

// AddArtical insert a new Artical into database and returns
// last inserted Id on success.
func AddLeave(m *Leave) (id int64, err error) {
	inserter, _ := orm.NewOrm().QueryTable(new(Leave)).PrepareInsert()
	o := orm.NewOrm()
	if id, err = o.Insert(m); err == nil {
		inserter.Close()
	}
	return id, err
}

func DeleteAttendance(id int) (err error) {
	o := orm.NewOrm()
	v := Attendance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Delete(&v); err == nil {
			return err
		}
	}
	return err
}

func GetAllLeave(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (a LeavePageList, err error) {
	var leavePageList LeavePageList
	o := orm.NewOrm()
	qs := o.QueryTable(new(Leave))
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
					return leavePageList, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return leavePageList, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return leavePageList, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return leavePageList, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Leave
	qs = qs.OrderBy(sortFields...)
	var ml []interface{}

	if count, err := qs.Count(); err == nil {
		leavePageList.TotalResult = count
	} else {
		leavePageList.TotalResult = 0
	}
	leavePageList.CurrentPage = offset
	leavePageList.ShowCount = limit

	offset = (offset - 1) * limit
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			if len(l) > 0 {
				for _, v := range l {
					m := make(map[string]interface{})
					val := reflect.ValueOf(v)
					for _, fname := range fields {
						m[fname] = val.FieldByName(fname).Interface()
					}
					ml = append(ml, m)
				}
			}
		}
		leavePageList.DataList = ml
		return leavePageList, nil
	}
	return leavePageList, err
}

func GetAllAttendance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (a AttendancePageList, err error) {
	var attendancePageList AttendancePageList
	o := orm.NewOrm()
	qs := o.QueryTable(new(Attendance))
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
					return attendancePageList, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return attendancePageList, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return attendancePageList, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return attendancePageList, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Attendance
	qs = qs.OrderBy(sortFields...)
	var ml []interface{}

	if count, err := qs.Count(); err == nil {
		attendancePageList.TotalResult = count
	} else {
		attendancePageList.TotalResult = 0
	}
	attendancePageList.CurrentPage = offset
	attendancePageList.ShowCount = limit

	offset = (offset - 1) * limit
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			if len(l) > 0 {
				for _, v := range l {
					m := make(map[string]interface{})
					val := reflect.ValueOf(v)
					for _, fname := range fields {
						m[fname] = val.FieldByName(fname).Interface()
					}
					ml = append(ml, m)
				}
			}
		}
		attendancePageList.DataList = ml
		return attendancePageList, nil
	}
	return attendancePageList, err
}

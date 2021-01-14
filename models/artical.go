package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)
type ArticalList struct{
	ShowCount int64 `json:"showCount"`
	CurrentPage int64 `json:"currentPage"`
	TotalResult int64 `json:"totalResult"`
	DataList []interface{} `json:"dataList"`
}

type TopAndNewArticalList struct{
	Id int64 `json:"id"`
	Title int64 `json:"title"`
	View int64 `json:"view"`
	CreateTime time.Time `json:"create_time"`
}

type Artical struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author	   string `json:"author"`
	View	   int `json:"view"`
	Content    string `json:"content"`
	CreateTime time.Time  `json:"create_time"          orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time  `json:"update_time"          orm:"auto_now";type(datetime)`
	RoleId     int64  `json:"role_id"`
	Tags       []*Tag `orm:"reverse(many)"`
}

type Tag struct {
	Id int `json:"id"`
	TagName string `json:"tag_name"`
	Artical *Artical `json:"-"                          orm:"rel(fk)"`
}

func init() {
	orm.RegisterModelWithPrefix("u_db_",new(Artical),new(Tag))
}

// AddArtical insert a new Artical into database and returns
// last inserted Id on success.
func AddArtical(m *Artical) (id int64, err error) {
	inserter, _ := orm.NewOrm().QueryTable(new(Tag)).PrepareInsert()
	o := orm.NewOrm()
	if id, err = o.Insert(m);err == nil {
		for _, v := range m.Tags {
			_,err = inserter.Insert(&Tag{TagName: v.TagName,Artical:&Artical{Id:int(id) }})
		}
		inserter.Close()
	}
	return id,err
}

// GetArticalById retrieves Artical by Id. Returns error if
// Id doesn't exist
func GetArticalById(id int) (v *Artical,err error) {
	orm.Debug = true
	o := orm.NewOrm()
	artical := Artical{Id: id}
	if  err = o.Read(&artical);err == nil {
		o.LoadRelated(&artical, "Tags")
		return &artical,nil
	}
	return nil,err
}


// GetArticalById retrieves Artical by Id. Returns error if
// Id doesn't exist
func GetArticalTags2ById(id int) (v Artical,err error) {
	o := orm.NewOrm()
	artical := Artical{Id: id}
	if err := o.Read(&artical); err == nil {
		o.LoadRelated(&artical, "Tag")
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
		articalList.TotalResult = count
	} else {
		articalList.TotalResult = 0
	}
	articalList.CurrentPage = offset
	articalList.ShowCount = limit

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
		articalList.DataList = ml
		return articalList, nil
	}
	return articalList, err
}

// UpdateArtical updates Artical by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticalById(id int,m *Artical) (err error) {
	o := orm.NewOrm()
	v := Artical{Id: id}
	if err = o.Read(&v); err == nil {
		if err  = UpdateTagsForArtical(id,m.Tags);err == nil {
			v.Title = m.Title
			v.Content = m.Content
			_, err = o.Update(&v)
		}
	}
	return err
}

/** Delete Articals Tag whith artical_id
 */
func UpdateTagsForArtical( id int ,tags []*Tag)(err error){
	orm.Debug = true
	ormer := orm.NewOrm().QueryTable(new(Tag))
	inserter, _ := ormer.PrepareInsert()
	if _, err := ormer.Filter("artical_id",id).Delete();err == nil {
			for _, v := range tags {
				_,err = inserter.Insert(&Tag{TagName: v.TagName,Artical:&Artical{Id:id }})
			}
	}
	return err
}




// DeleteArtical deletes Artical by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArtical(id int) (err error) {
	o := orm.NewOrm()
	v := Artical{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		o.LoadRelated(&v, "Tags")
		if _, err = o.Delete(&v); err == nil {
			return  err
		}
	}
	return err
}
//get Top and new artical for home/index
func GetTopAndNewArticalList( size int64 )(list map[string]interface{},err error) {
	var TopList []*Artical
	o := orm.NewOrm()
	qs := o.QueryTable(new(Artical))
	num ,err := qs.OrderBy("-view").Limit(size).All(&TopList,"Id","View","Title","CreateTime")
	fmt.Println(num,err,TopList)
	if num == 0 || err != nil{
		return nil,err
	}
	var  NewList []*Artical
	num , err  = qs.OrderBy("-create_time").Limit(size).All(&NewList,"Id","View","Title","CreateTime")
	if num == 0 ||  err != nil{
		return nil, err
	}
	list["TopList"] = {"fdas","fdasf"}
	list["NewList"] = string{"1","2"}
	return list,nil
}






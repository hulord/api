package utils

import (
	"errors"
	"reflect"
	"regexp"
	// "fmt"
	//"api/models"
)

func main(){
	
}
type checkParam struct {

}

type TreeList struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	ParentId   int    `json:"parent_id"`
	Path       string `json:"path"`
	CreateTime int    `json:"create_time"`
    Children []TreeList	`json:"children"`
}

//元素是否中数组中
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		match,_:=regexp.MatchString(eachItem,item)
		if match {
			return true
		}
	}
	return false
}

//切片中是否存在空值并返回
func IsEmpty(params map[string]interface{},limitMaps []string) (err error) {
	 if len(params) != 0 {
		 for _,val := range limitMaps{
		 	if params[val]  == "" {
				return errors.New(val+"不能为空")
			}
		 }
		 return nil
	 }else{
		 return errors.New("数据包不能为空")
	 }
}

//数组扁平化
func Tree(treeMap []map[string]interface{},pid int)[]TreeList{
	branch := make([]TreeList, 0)
	 for j := 0; j < len(treeMap); j++ {
		if int(pid) == treeMap[j]["ParentId"].(int) {
			child := TreeList{
				Id :		  treeMap[j]["Id"].(int),
				Name:         treeMap[j]["Name"].(string),
				Path:         treeMap[j]["Path"].(string),
				Icon:         treeMap[j]["Icon"].(string),
				Children: Tree(treeMap, treeMap[j]["Id"].(int)),
			}

			branch = append(branch, child)
		}
	 }
	return branch
}

//struct转map
func StructToMapDemo(obj interface{}) map[string]interface{}{
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}






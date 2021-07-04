package utils

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strconv"
	// "fmt"
	//"api/models"
)

func main() {

}

type checkParam struct {
}

type TreeList struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Icon       string     `json:"icon"`
	ParentId   int        `json:"parent_id"`
	Path       string     `json:"path"`
	CreateTime int        `json:"create_time"`
	Children   []TreeList `json:"children"`
}

//元素是否中数组中
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		match, _ := regexp.MatchString(eachItem, item)
		if match {
			return true
		}
	}
	return false
}

//切片中是否存在空值并返回
func IsEmpty(params map[string]interface{}, limitMaps []string) (err error) {
	if len(params) != 0 {
		for _, val := range limitMaps {
			if params[val] == "" {
				return errors.New(val + "不能为空")
			}
		}
		return nil
	} else {
		return errors.New("数据包不能为空")
	}
}

//数组扁平化
func Tree(treeMap []map[string]interface{}, pid int) []TreeList {
	branch := make([]TreeList, 0)
	for j := 0; j < len(treeMap); j++ {
		if int(pid) == treeMap[j]["ParentId"].(int) {
			child := TreeList{
				Id:       treeMap[j]["Id"].(int),
				Name:     treeMap[j]["Name"].(string),
				Path:     treeMap[j]["Path"].(string),
				Icon:     treeMap[j]["Icon"].(string),
				Children: Tree(treeMap, treeMap[j]["Id"].(int)),
			}

			branch = append(branch, child)
		}
	}
	return branch
}

//struct转map
func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// json转map函数，通用
func JSONToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}

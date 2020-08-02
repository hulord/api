package utils

import (
	"regexp"
	// "fmt"
	//"api/models"

)

func main(){
	
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

//tree
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




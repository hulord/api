package utils

import (
	"regexp"
	"fmt"
)

func main(){
	
}

//元素是否中数组中
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		match,_:=regexp.MatchString(eachItem,item)
		fmt.Println(match)
		if match {
			return true
		}
	}
	return false

	// for _, eachItem := range items {
	// 	if eachItem == item {
	// 		return true
	// 	}
	// }
	// return false
}


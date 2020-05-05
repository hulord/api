package utils

import (

)

func main(){
	
}

//元素是否中数组中
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}


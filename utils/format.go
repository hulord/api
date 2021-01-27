package utils

import (
	"encoding/json"
)




/**
	数据类型转换类
 */
func  Format(objective string,output string){
	switch objective {
		case "":

	}
}
/**
	字符串转数字 string=>int
 */
func FormatStringToInt(string string,int int){

}

func FormatStringToJson( string string )( ret map[string]interface{},err error ){
	if err := json.Unmarshal ([]byte(string), &ret);err != nil{
		return nil,err
	}
	return ret,nil
}










package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	
}

type Tag struct {
	Id       int 		`json:"id"`
	Type string		`json:"type"`
	TagName string		`json:"tagName"`
}


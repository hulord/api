package models

import (
)

func init() {
	
}

type Tag struct {
	Id      int      `json:"id"`
	Type    string	 	`json:"type"`
	TagName string   `json:"tag_name"`
	Artical *Artical `json:"-" orm:"rel(fk)";on_delete(set_null)"`
}


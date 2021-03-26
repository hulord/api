package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Image_20210326_162521 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Image_20210326_162521{}
	m.Created = "20210326_162521"

	migration.Register("Image_20210326_162521", m)
}

// Run the migrations
func (m *Image_20210326_162521) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE image(`id` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL,`type` varchar(128) NOT NULL,`url` varchar(128) NOT NULL,`create_time` int(11) DEFAULT NULL,`update_time` int(11) DEFAULT NULL)")
}

// Reverse the migrations
func (m *Image_20210326_162521) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `image`")
}

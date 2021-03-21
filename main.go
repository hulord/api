package main

import (
	_ "api/routers"
	"api/models"
	"api/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//注册 model
	orm.RegisterModelWithPrefix("u_db_",new(models.User))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "root:root@(localhost:3306)/test?charset=utf8")//密码为空格式
	orm.RegisterDataBase("default", "mysql", "root:19930202a@(cdb-om56p9rc.cd.tencentcdb.com:10014)/u_db_site?charset=utf8")//密码为空格式
	orm.RunSyncdb("default", false, false)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}



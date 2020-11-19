package controllers
 

import (
	"github.com/astaxie/beego"
)

// ErrorController operations for Error
type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.ApiJsonReturn(404,"非法路由","")	
}
func (c *ErrorController) Error500() {
	c.Abort("500")
	c.ApiJsonReturn(500,beego.Error("错误"),"")	
}

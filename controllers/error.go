package controllers
 

import (
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
	c.ApiJsonReturn(500,"未知错误","")
}

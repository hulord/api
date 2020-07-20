package controllers

import (
)

// MenuController operations for Menu
type MenuController struct {
	BaseController
}

// URLMapping ...
func (c *MenuController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post
// @Description asking Menu
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 201 {map} Menu.list
// @Failure 403 body is empty
// @router / [post]
func (c *MenuController) GetAll() {
	c.ApiJsonReturn(0,"MSG","1")	
}


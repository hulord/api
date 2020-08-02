package controllers

import (
	//"fmt"
	// "encoding/json"
	"api/models"
	utils "api/utils"
	//"reflect"
)

// MenuController operations for Menu
type MenuController struct {	
	BaseController
}
type Menus struct {
    Name              string   `json:"name"`
    Icon              string    `json:"icon"`
	Path     		  string   `json:"path"`
	Children		  []*Menu   `json:"children"`
}

type Menu struct {
    Name              string   `json:"name"`
    Icon              string    `json:"icon"`
	Path     		  string   `json:"path"`
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
	// jsonStr := `{"name":"仪表盘","icon":"dashboard","path": "/dashboard","children":{"name":"仪表盘1","icon":"dashboard","path": "/dashboard"}}`
	// var dat  
	// err := json.Unmarshal([]byte(jsonStr), &dat)
	// fmt.Println(dat,err)

	// var menus Menus
	// err := json.Unmarshal([]byte(jsonStr), &menus)
	// fmt.Println(menus,err)
	m, err := models.GetMenuByRole(c.role)
	//jsons,_ := json.Marshal(m)
	if  err != nil {
		c.ApiJsonReturn(1,err.Error(),"")	
	} else {		
		// for _, var_val := range m {
		// 	fmt.Println(var_val.Menu.ParentId)
		// }
		menu := utils.Tree(m,1)
		c.ApiJsonReturn(0,"",menu)	
	}
	c.ApiJsonReturn(0,"fsda","1")	
}


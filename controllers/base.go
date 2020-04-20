package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type JsonReturn struct {
	Status int		    `json:"status"`
	Message  string 	`json:"message"`
	Data interface{}	`json:"data"`		//Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`	
}

func (c *BaseController) ApiJsonReturn(status int,message string,data interface{}) {
	var JsonReturn JsonReturn
	JsonReturn.Status = status
	JsonReturn.Message = message
	JsonReturn.Data = data
	c.Data["json"] = JsonReturn		//将结构体数组根据tag解析为json
	c.ServeJSON()					//对json进行序列化输出
	c.StopRun()						//终止执行逻辑
}

func (b *BaseController) Oauth(){

}

func (b *BaseController) GetUserInfo(){

}

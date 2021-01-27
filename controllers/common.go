package controllers

import (
	"api/utils"
)

// CommonController operations for common
type CommonlController struct {
	BaseController
}

// getWeather ...
// @Title getWeather
// @Description get Artical by id
// @Param	cityId		path 	string	true		"The key for staticblock"
// @Success 200 {object}
// @Failure 403 :cityId is empty
// @router /getWeather/:cityId [get]
func (c *CommonlController) GetWeather() {
	getWeatherUrl := "http://t.weather.itboy.net/api/weather/city/"
	cityId := c.Ctx.Input.Param(":cityId")
	if len(cityId) == 0 {
		c.ApiJsonReturn(1,"未指定具体城市","")
	}
	res, err := utils.Get(getWeatherUrl+cityId)
	if err != nil {
		c.ApiJsonReturn(1,err.Error(),"")
	}else {
		if ret,err := utils.FormatStringToJson(res);err != nil{
			c.ApiJsonReturn(1,err.Error(),"")
		}else{
			c.ApiJsonReturn(0,"1",ret)
		}
	}
}

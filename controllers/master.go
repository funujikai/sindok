package controllers

import (
	"PMM/models"
	"PMM/global"

	// "encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// Endpoint
type MasterController struct {
	beego.Controller
}

// @Title API Data Nasabah
// @Description Get Data Nasabah Mekaar
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /cabang [get]
func (c *MasterController) Cabang() {

	data,err := models.Cabang()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = data
	c.ServeJSON()
}
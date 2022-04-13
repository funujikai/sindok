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

// @Title master cabang
// @Description get data cabang
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {string} models.Master_cabang
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

// @Title master tipe file
// @Description get data tipe file
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {string} models.Master_tipefile
// @Failure 403 body is empty
// @router /tipe-file [get]
func (c *MasterController) Tipe_file() {

	data,err := models.Tipefile()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = data
	c.ServeJSON()
}

// @Title master role
// @Description get data role
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {string} models.Master_role
// @Failure 403 body is empty
// @router /role [get]
func (c *MasterController) Role() {

	data,err := models.Role()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = data
	c.ServeJSON()
}


// @Title master role detail
// @Description get data role detail
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {string} models.Master_role_detail
// @Failure 403 body is empty
// @router /role-detail [get]
func (c *MasterController) Role_detail() {

	data,err := models.RoleDetail()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = data
	c.ServeJSON()
}
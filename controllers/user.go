package controllers

import (
	"PMM/models"
	"PMM/global"
	// "encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// Endpoint
type UserController struct {
	beego.Controller
}


// @Title Login Cabang
// @Description Login menggunakan User Mekaar Integrasi
// @Param	User formData string  false "Username ex: 90004"
// @Param	Pass formData string  false "Password ex: 12345678"
// @Success 200 {string} models.T_mkr_users
// @Failure 403 body is empty
// @router /login-cabang [post]
func (c *UserController) LoginCabang() {
	user := c.GetString("User")
	password := c.GetString("Pass")

	data,err := models.Login_mekaar_integrasi(user,password)	
	if err != nil {		
		// global.Logging("ERROR","controllers.Login ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}	

	c.Data["json"] = global.APIDataResponse{Code: 200,Message: "Berhasil Login",Data: data}
	c.ServeJSON()
}


// @Title Login Pusat
// @Description Login menggunakan User SSO
// @Param	User formData string  false "Username ex: WALubis0508"
// @Param	Pass formData string  false "Password ex: pnm12345"
// @Success 200 {string} models.T_mkr_users
// @Failure 403 body is empty
// @router /login-pusat [post]
func (c *UserController) LoginPusat() {
	user := c.GetString("User")
	password := c.GetString("Pass")

	data,err := models.Login_sso(user,password)	
	if err != nil {		
		// global.Logging("ERROR","controllers.Login ---> " + err.Error())		
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}	

	c.Data["json"] = global.APIDataResponse{Code: 200,Message: "Berhasil Login",Data: data}
	c.ServeJSON()
}

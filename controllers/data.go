package controllers

import (
	"PMM/models"
	"PMM/global"

	// "encoding/json"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

// Endpoint
type DataController struct {
	beego.Controller
}

// @Title API Data Nasabah
// @Description Get Data Nasabah Mekaar
// @Param	Authorization header string  false "Authorization Token"
// @Param	cabangid path string  true "CabangID ex: 90001"
// @Param	kelompokid path string  false "KelompokID ex: 90001007  isi "null" jika tanpa kelompokid"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /nasabah/:cabangid/:kelompokid [get]
func (c *DataController) Nasabah() {

	CabangID := c.Ctx.Input.Param(":cabangid")
	KelompokID := c.Ctx.Input.Param(":kelompokid")

	data,err := models.Nasabah(CabangID,KelompokID)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = data
	c.ServeJSON()
}

// @Title API Data Kelompok
// @Description Get Data Kelompok Mekaar
// @Param	Authorization header string  false "Authorization Token"
// @Param	cabangid path string  true "CabangID ex: 90108"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /kelompok/:cabangid [get]
func (c *DataController) Kelompok() {

	CabangID := c.Ctx.Input.Param(":cabangid")

	data,err := models.Kelompok(CabangID)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = data
	c.ServeJSON()
}

// @Title API Data Kelompok
// @Description Get Data Kelompok Mekaar
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body models.Data_set_upload_file true "request data untuk di update"
// @Success 200 {object} models.Data_set_upload_file
// @Failure 403 body is empty
// @router /set-upload-file [post]
func (c *DataController) SetUploadFile() {

	var data models.Data_set_upload_file
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	

	res,err := models.SetUploadFile(data)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}

	// c.Data["json"] = global.APIResponse{Code: 200,Message: "Update Berhasil"}
	// c.ServeJSON()

	
	c.Data["json"] = global.APIDataResponse{Code: 200,Message: "Berhasil Login",Data: res}
	c.ServeJSON()
}
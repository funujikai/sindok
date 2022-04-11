package controllers

import (
	// "PMM/models"
	"PMM/global"

	// "encoding/json"
	// "encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"io"
)

// Endpoint
type OtherController struct {
	beego.Controller
}

// @Title PkmImages
// @Description PkmImages
// @Param   data path string true "file"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /file/:data [get]
func (c *OtherController) ShowS3Image() {	

	object,err := global.GetS3(c.Ctx.Input.Param(":data"))
	if err != nil {
		// global.Logging("ERROR","global.ConnS3Storage controller ShowS3Image ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }    


	c.Ctx.Output.Header("Content-Type", "image/png")
	if _, err := io.Copy(c.Ctx.ResponseWriter,object); err != nil {
		// global.Logging("ERROR","io.Copy controller ShowS3Image ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}				
	
}
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
// @Param   tipe path string true "tipe file ex images | pdf"
// @Success 200 {object} global.APIResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /file/:data/:tipe [get]
func (c *OtherController) ShowS3Image() {	

	object,err := global.GetS3(c.Ctx.Input.Param(":data"))
	if err != nil {
		// global.Logging("ERROR","global.ConnS3Storage controller ShowS3Image ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }    

	switch c.Ctx.Input.Param(":tipe"){
	case "images":
		c.Ctx.Output.Header("Content-Type", "image/png")
		if _, err := io.Copy(c.Ctx.ResponseWriter,object); err != nil {
			// global.Logging("ERROR","io.Copy controller ShowS3Image ---> " + err.Error())	
			c.Ctx.Output.SetStatus(500)	
			c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
			c.ServeJSON()
		}			
	case "pdf":
		c.Ctx.Output.Header("Content-Type", "application/pdf")	
		c.Data["json"] = object
		c.ServeJSON()			
	default:
		c.Data["json"] = global.APIResponse{Code: 500, Message: "tipe file tidak di izinkan"}
		c.ServeJSON()
	}

	
}
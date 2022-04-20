package controllers

import (
	"PMM/models"
	"PMM/global"
	"fmt"
	"os"
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
func (c *OtherController) ShowS3File() {	

	object,err := global.GetS3(c.Ctx.Input.Param(":data"))
	if err != nil {
		// global.Logging("ERROR","global.ConnS3Storage controller ShowS3Image ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }    

	
	c.Ctx.Output.Header("Content-type", "application/octet-stream")
	if _, err := io.Copy(c.Ctx.ResponseWriter,object); err != nil {
		// global.Logging("ERROR","io.Copy controller ShowS3Image ---> " + err.Error())	
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}				


	// f, err := os.Open("./temp/"+c.Ctx.Input.Param(":data"))
    // if err != nil {
	// 	c.Ctx.Output.SetStatus(500)	
	// 	c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
	// 	c.ServeJSON()
    // }
    // defer f.Close()		

	// c.Ctx.Output.Header("Content-Type", "application/pdf")	
	// contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	// c.Ctx.Output.Header("Content-Disposition", contentDisposition)

	// if _, err := io.Copy(c.Ctx.ResponseWriter, f); err != nil {
	// 	c.Ctx.Output.SetStatus(500)	
	// 	c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
	// 	c.ServeJSON()
    // }
	
}

// @Title PkmImages
// @Description PkmImages
// @Param	cabangid path string  true "CabangID ex: 90001"
// @Param	loanid path string  true "Loanid ex: 9000101M2000002"
// @Param	siklus path string  true "Siklus ex: 1"
// @Success 200 "file zip"
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql"}
// @router /download-zip-file/:cabangid/:loanid/:siklus [get]
func (c *OtherController) DownloadZipS3File() {

	cabangid := c.Ctx.Input.Param(":cabangid")
	loanid := c.Ctx.Input.Param(":loanid")
	siklus := c.Ctx.Input.Param(":siklus")


	err := models.SetDownloadFile(cabangid,loanid,siklus)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIResponse{Code: 500,Message: err.Error()}
		c.ServeJSON()
	}	

	

	f, err := os.Open("./temp/data"+cabangid+loanid+siklus+".zip")
    if err != nil {
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }
    defer f.Close()		

	c.Ctx.Output.Header("Content-type", "application/octet-stream")
	// contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	contentDisposition := fmt.Sprintf("attachment; filename=%s", cabangid+"-"+loanid+"-"+siklus+".zip")
	c.Ctx.Output.Header("Content-Disposition", contentDisposition)
	if _, err := io.Copy(c.Ctx.ResponseWriter, f); err != nil {
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APIResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
    }else{
		os.RemoveAll("./temp/")
	}
}
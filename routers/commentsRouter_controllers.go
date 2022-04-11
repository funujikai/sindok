package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["PMM/controllers:DataController"] = append(beego.GlobalControllerRouter["PMM/controllers:DataController"],
        beego.ControllerComments{
            Method: "Kelompok",
            Router: "/kelompok/:cabangid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["PMM/controllers:DataController"] = append(beego.GlobalControllerRouter["PMM/controllers:DataController"],
        beego.ControllerComments{
            Method: "Nasabah",
            Router: "/nasabah/:cabangid/:kelompokid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["PMM/controllers:DataController"] = append(beego.GlobalControllerRouter["PMM/controllers:DataController"],
        beego.ControllerComments{
            Method: "SetUploadFile",
            Router: "/set-upload-file",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["PMM/controllers:MasterController"] = append(beego.GlobalControllerRouter["PMM/controllers:MasterController"],
        beego.ControllerComments{
            Method: "Cabang",
            Router: "/cabang",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["PMM/controllers:OtherController"] = append(beego.GlobalControllerRouter["PMM/controllers:OtherController"],
        beego.ControllerComments{
            Method: "ShowS3Image",
            Router: "/file/:data/:tipe",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["PMM/controllers:UserController"] = append(beego.GlobalControllerRouter["PMM/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginCabang",
            Router: "/login-cabang",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["PMM/controllers:UserController"] = append(beego.GlobalControllerRouter["PMM/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginPusat",
            Router: "/login-pusat",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

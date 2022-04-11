// @APIVersion 1.0.0
// @Title API APLIKASI PMM
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"PMM/controllers"
	"PMM/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(middleware.Jwt),
		beego.NSNamespace("/data",
			beego.NSInclude(
				&controllers.DataController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/master",
			beego.NSInclude(
				&controllers.MasterController{},
			),
		),		
		beego.NSNamespace("/other",
			beego.NSInclude(
				&controllers.OtherController{},
			),
		),		

	)
	beego.AddNamespace(ns)
}

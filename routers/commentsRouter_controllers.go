package routers

import "github.com/astaxie/beego"

func init() {

	beego.GlobalControllerRouter["restful_api/controllers : LoginController"] = append(beego.GlobalControllerRouter["restful_api/controllers : LoginController"],
		beego.ControllerComments{
			Method:"Get",
			Router: `/:username/:passwd`,
			AllowHTTPMethods: []string{"get"},
			Params: nil,
		})

	beego.GlobalControllerRouter["restful_api/controllers : LoginController"] = append(beego.GlobalControllerRouter["restful_api/controllers : LoginController"],
		beego.ControllerComments{
			Method:"Post",
			Router: `/`,
			AllowHTTPMethods: []string{"Post"},
			Params: nil,
		})

}
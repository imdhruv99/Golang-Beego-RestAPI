package routers

import (
	"github.com/astaxie/beego"
	"restful_api/controllers"
)

func init() {
	var namespaces []string = []string {"MyAPI", "another"}

	var controller []string = []string{"login", "other"}

	restfulRouter := beego.NewNamespace("/" + namespaces[0],
		beego.NSNamespace("/" + controller[0],
			beego.NSInclude(&controllers.LoginController{},
				),
			),
		)
	beego.AddNamespace(restfulRouter)
}

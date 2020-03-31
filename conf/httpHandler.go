package conf

import (
	"github.com/astaxie/beego"
)

func RestfulAPIServiceInit(method string){
	beego.BConfig.RunMode = "dev" 
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.EnableDocs = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.Listen.HTTPSCertFile = "tls-ssl/file-rest.crt"
	beego.BConfig.Listen.HTTPSKeyFile = "tls-ssl/file-rest.key"
	beego.BConfig.Listen.HTTPPort = 8080
	beego.BConfig.Listen.HTTPSPort = 8081
	if method == "HTTP" {
	  beego.BConfig.Listen.EnableHTTP = true
	  beego.BConfig.Listen.EnableHTTPS = false
	} else if method == "HTTPS" {
	  beego.BConfig.Listen.EnableHTTP = false
	  beego.BConfig.Listen.EnableHTTPS = true
	}
	  beego.Run()
  }
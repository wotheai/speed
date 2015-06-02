package routers

import (
	"github.com/astaxie/beego"
	"os"
	"speed/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	os.Mkdir("attach", os.ModePerm)
	beego.Router("/img/:all", &controllers.ImgController{})
	beego.Router("/attach/", &controllers.AttachController{})
	beego.AutoRouter(&controllers.AttachController{})
	beego.Router("/json/", &controllers.JsonController{})
}

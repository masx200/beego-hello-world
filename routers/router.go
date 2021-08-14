package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	controllers "hello/controllers"
)

func Init() {
	beego.Router("/beego", &controllers.MainController{})
}

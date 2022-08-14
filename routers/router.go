package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	controllers "github.com/masx200/beego-hello-world/controllers"
)

func Init() {
	beego.Router("/beego", &controllers.MainController{})
}

package helloworld

import web "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
}

func Helloworld() {
	web.Router("/helloworld", &MainController{})
	// web.Run()
}

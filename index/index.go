package index

import (
	beego "github.com/beego/beego/v2/server/web"
)

func Index() {
	beego.Router("/", &MainController{})

}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplName = "index.html"
}

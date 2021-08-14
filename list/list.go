package list

import (
	// "fmt"

	web "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}
type A struct {
	Name string `json:"name"`
}

func (c *MainController) Get() {
	a := []A{{Name: "golang"}, {Name: "python"}, {Name: "html"}, {Name: "css"}}
	// fmt.Printf("%#v",a)
	c.Data["json"] = a
	c.ServeJSON()
}

func List() {
	web.Router("/list", &MainController{})
	// web.Run()
}

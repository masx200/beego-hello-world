package main

import (
	beego "github.com/beego/beego/v2/server/web"

	helloworld "hello/helloworld"
	index "hello/index"
	list "hello/list"
	routers "hello/routers"
)

func main() {
	list.List()
	routers.Init()
	helloworld.Helloworld()
	index.Index()
	beego.Run()
}

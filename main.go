package main

import (
	beego "github.com/beego/beego/v2/server/web"

	cors "github.com/beego/beego/v2/server/web/filter/cors"
	helloworld "github.com/masx200/beego-hello-world/helloworld"
	index "github.com/masx200/beego-hello-world/index"
	list "github.com/masx200/beego-hello-world/list"
	routers "github.com/masx200/beego-hello-world/routers"
)

func addcors() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
func main() {
	addcors()
	list.List()
	routers.Init()
	helloworld.Helloworld()
	index.Index()
	beego.Run("0.0.0.0:9000")
}

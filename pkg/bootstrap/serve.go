package bootstrap

import (
	"goblog/pkg/config"
	"goblog/pkg/database"
	"goblog/pkg/html"
	"goblog/pkg/routing"
	"goblog/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	routing.RegisterRoutes()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTLM(routing.GetRouter())

	routing.Serve()
}

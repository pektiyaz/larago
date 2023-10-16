package routes

import (
	"larago/routes"
	index_controller "larago/src/app/Http/Controllers"
)

func Web() *routes.Route {
	route := routes.New()

	route.Get("/", index_controller.Index)
	route.Get("/json", index_controller.JsonTest)

	return route

}

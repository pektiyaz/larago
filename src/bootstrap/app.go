package bootstrap

import (
	"larago/src/routes"
	"larago/webserver"
)

func New() {
	route := routes.Web()
	webserver.Serve(route)
}

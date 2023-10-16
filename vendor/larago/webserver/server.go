package webserver

import (
	"larago/routes"
	"log"
	"net/http"
)

func Serve(router *routes.Route) {
	log.Fatal(http.ListenAndServe(":8080", router.Route))
}

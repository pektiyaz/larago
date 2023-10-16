package routes

import (
	"julienschmidt/httprouter"
	"net/http"
)

type Request struct {
	Req      *http.Request
	Params   httprouter.Params
	Response http.ResponseWriter
}

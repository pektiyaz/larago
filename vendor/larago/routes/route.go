package routes

import (
	"julienschmidt/httprouter"
	"net/http"
)

type Route struct {
	Route *httprouter.Router
}

func New() *Route {
	return &Route{
		Route: httprouter.New(),
	}

}

func (r *Route) Get(url string, callback func(request Request) interface{}) {

	r.Route.GET(url, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handle(callback, w, r, ps)
	})
}
func (r *Route) Post(url string, callback func(request Request) interface{}) {

	r.Route.POST(url, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handle(callback, w, r, ps)
	})
}

func (r *Route) Put(url string, callback func(request Request) interface{}) {

	r.Route.PUT(url, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handle(callback, w, r, ps)
	})
}
func (r *Route) Patch(url string, callback func(request Request) interface{}) {

	r.Route.PATCH(url, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handle(callback, w, r, ps)
	})
}
func (r *Route) Delete(url string, callback func(request Request) interface{}) {

	r.Route.DELETE(url, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handle(callback, w, r, ps)
	})
}

func handle(callback func(request Request) interface{}, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	request := Request{
		Req:      r,
		Params:   ps,
		Response: w,
	}

	response := callback(request)

	switch response.(type) {
	case []byte:
		_, err := w.Write(response.([]byte))
		if err != nil {
			panic(err)
		}
	case string:
		_, err := w.Write([]byte(response.(string)))
		if err != nil {
			panic(err)
		}
	}

}

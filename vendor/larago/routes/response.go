package routes

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Body http.ResponseWriter
}

func JsonResponse(v any) string {
	val, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(val)
}

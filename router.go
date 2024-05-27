package myrouter

import (
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request, Params)

type Params map[string]string

type Route struct {
	method  string
	path    string
	handler Handler
}

type Router struct {
	routes []Route
}

func New() *Router {
	return &Router{}
}

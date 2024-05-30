package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"
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

type contextKey struct{}

func WithParams(r *http.Request, params Params) *http.Request {
	ctx := context.WithValue(r.Context(), contextKey{}, params)
	return r.WithContext(ctx)
}

func GetParams(r *http.Request) Params {
	params, ok := r.Context().Value(contextKey{}).(Params)
	if ok {
		return params
	}
	return Params{}
}

func (r *Router) Handle(method, path string, handler Handler) {
	r.routes = append(r.routes, Route{method, path, handler})
}

func (r *Router) GET(path string, handler Handler) {
	r.Handle("GET", path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	r.Handle("POST", path, handler)
}

// ServeHTTP implement the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if match, params := r.Match(route, req); match {
			req := WithParams(req, params)
			route.handler(w, req, params)
			return
		}
	}
	fmt.Println(r.routes)

	http.NotFound(w, req)
}

func (m *Router) Match(route Route, req *http.Request) (bool, Params) {
	if route.method != req.Method {
		return false, nil
	}

	routeParts := strings.Split(route.path, "/")
	reqParts := strings.Split(req.URL.Path, "/")

	fmt.Println("routePart=", routeParts)
	fmt.Println("reqParts=", reqParts)

	if len(routeParts) != len(reqParts) {
		return false, nil
	}

	params := Params{}

	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], ":") {
			params[routeParts[i][1:]] = reqParts[i]
		} else if routeParts[i] != reqParts[i] {
			return false, nil
		}
	}

	fmt.Println("params", params["id"])
	return true, params

}

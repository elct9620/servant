package router

import "net/http"

type localRouter struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func (r *localRouter) Path() string {
	return r.path
}

func (r *localRouter) Method() string {
	return r.method
}

func (r *localRouter) Handler() http.HandlerFunc {
	return r.handler
}

func NewGetRoute(path string, handler http.HandlerFunc) Route {
	return &localRouter{
		method:  http.MethodGet,
		path:    path,
		handler: handler,
	}
}

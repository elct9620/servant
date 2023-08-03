package router

import (
	"net/http"

	"github.com/elct9620/servant/api/server/httputils"
)

type localRouter struct {
	method  string
	path    string
	handler httputils.ApiFunc
}

func (r *localRouter) Path() string {
	return r.path
}

func (r *localRouter) Method() string {
	return r.method
}

func (r *localRouter) Handler() httputils.ApiFunc {
	return r.handler
}

func NewGetRoute(path string, handler httputils.ApiFunc) Route {
	return &localRouter{
		method:  http.MethodGet,
		path:    path,
		handler: handler,
	}
}

package router

import "github.com/elct9620/servant/api/server/httputils"

type Router interface {
	Routes() []Route
}

type Route interface {
	Path() string
	Method() string
	Handler() httputils.ApiFunc
}

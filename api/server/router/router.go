package router

import "net/http"

type Router interface {
	Routes() []Route
}

type Route interface {
	Path() string
	Method() string
	Handler() http.HandlerFunc
}

package server

import (
	"net/http"

	"github.com/elct9620/servant/api/server/router"
	"github.com/go-chi/chi/v5"
)

const versionMatcher = "/v{version:[0-9]+}"

func New(routers ...router.Router) http.Handler {
	mux := chi.NewRouter()

	for _, router := range routers {
		for _, route := range router.Routes() {
			mux.Method(route.Method(), versionMatcher+route.Path(), route.Handler())
			mux.MethodFunc(route.Method(), route.Path(), route.Handler())
		}
	}

	return mux
}

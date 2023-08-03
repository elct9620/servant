package server

import (
	"net/http"

	"github.com/elct9620/servant/api/server/httputils"
	"github.com/elct9620/servant/api/server/router"
	"github.com/go-chi/chi/v5"
)

const versionMatcher = "/v{version:[0-9]+}"

type ErrorResponse struct {
	Message string `json:"message"`
}

func New(routers ...router.Router) http.Handler {
	mux := chi.NewRouter()

	for _, router := range routers {
		for _, route := range router.Routes() {
			handler := convertHttpHandler(route.Handler())

			mux.Method(route.Method(), versionMatcher+route.Path(), handler)
			mux.MethodFunc(route.Method(), route.Path(), handler)
		}
	}

	return mux
}

func newHttpErrorHandler(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = httputils.WriteJson(w, http.StatusNotFound, &ErrorResponse{
			Message: err.Error(),
		})
	}
}

func convertHttpHandler(handler httputils.ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			newHttpErrorHandler(err)(w, r)
		}
	}
}

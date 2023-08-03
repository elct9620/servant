package system

import (
	"net/http"

	"github.com/elct9620/servant/api/server/router"
)

type SystemRouter struct {
	routes []router.Route
}

func New() *SystemRouter {
	r := &SystemRouter{}

	r.routes = []router.Route{
		router.NewGetRoute("/livez", r.Liveness),
		router.NewGetRoute("/readz", r.Readiness),
	}

	return r
}

func (r *SystemRouter) Routes() []router.Route {
	return r.routes
}

func (sys *SystemRouter) Liveness(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)

	return nil
}

func (sys *SystemRouter) Readiness(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)

	return nil
}

package servant

import (
	"net/http"

	"github.com/elct9620/servant/api/server"
	"github.com/elct9620/servant/api/server/router/system"
)

func NewApi() http.Handler {
	return server.New(
		system.New(),
	)
}

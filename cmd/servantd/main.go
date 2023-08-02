package main

import (
	"net/http"

	"github.com/elct9620/servant/api/server"
	"github.com/elct9620/servant/api/server/router/system"
)

func main() {
	server := server.New(
		system.New(),
	)

	http.ListenAndServe(":8080", server)
}

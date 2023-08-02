package main

import (
	"net/http"

	"github.com/elct9620/servant/api/server"
	"github.com/elct9620/servant/api/server/router/system"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	server := server.New(
		system.New(),
	)

	err = http.ListenAndServe(":8080", server)
	if err != nil {
		logger.Fatal("unable to start server", zap.Error(err))
	}
}

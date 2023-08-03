package main

import (
	"net/http"

	"github.com/elct9620/servant"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	api := servant.NewApi()

	err = http.ListenAndServe(":8080", api)
	if err != nil {
		logger.Fatal("unable to start server", zap.Error(err))
	}
}

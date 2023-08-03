package daemon

import (
	"context"
	"errors"
	"net/http"
)

type HttpService struct {
	http.Server
}

func NewHttpService(handler http.Handler) *HttpService {
	return &HttpService{
		Server: http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
	}
}

func (svc *HttpService) Start(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		svc.Stop(ctx)
	}()

	err := svc.Server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (svc *HttpService) Stop(ctx context.Context) error {
	return svc.Server.Shutdown(ctx)
}

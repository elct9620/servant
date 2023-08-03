package daemon

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Daemon struct {
	services []Service
}

func New(services ...Service) *Daemon {
	return &Daemon{
		services: services,
	}
}

func (d *Daemon) Run(ctx context.Context) error {
	runner, runCtx := errgroup.WithContext(ctx)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	for _, service := range d.services {
		service := service
		runner.Go(func() error {
			return service.Start(runCtx)
		})
	}

	go func() {
		<-exit
		for _, service := range d.services {
			service := service

			runner.Go(func() error {
				return service.Stop(runCtx)
			})
		}
	}()

	return runner.Wait()
}

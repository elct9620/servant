package watcher

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const exitNormal = 0

type Watcher struct {
	client client.APIClient
	exit   chan int
}

func New(client client.APIClient) *Watcher {
	return &Watcher{
		client: client,
		exit:   make(chan int, 1),
	}
}

func (w *Watcher) Start(ctx context.Context) error {
	events, errors := w.client.Events(ctx, types.EventsOptions{
		Filters: filters.NewArgs(),
	})

	for {
		select {
		case <-events:
		case err := <-errors:
			if err == io.EOF {
				return nil
			}

			return err
		case <-ctx.Done():
			return nil
		case <-w.exit:
			return nil
		}
	}
}

func (w *Watcher) Stop(ctx context.Context) error {
	w.exit <- exitNormal
	return nil
}

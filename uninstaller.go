package servant

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Uninstaller struct {
}

func (u *Uninstaller) Execute(ctx context.Context, api client.APIClient) error {
	fmt.Println("Uninstalling the servant controller...")
	err := u.UninstallController(ctx, api)
	if err != nil {
		return err
	}

	fmt.Println("Uninstalling the servant network...")
	err = u.UninstallNetwork(ctx, api)
	if err != nil {
		return err
	}

	return nil
}

func (u *Uninstaller) UninstallController(ctx context.Context, api client.ServiceAPIClient) error {
	services, err := api.ServiceList(ctx, types.ServiceListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", ControllerName),
		),
	})

	if err != nil {
		return err
	}

	for _, service := range services {
		err = api.ServiceRemove(ctx, service.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *Uninstaller) UninstallNetwork(ctx context.Context, api client.NetworkAPIClient) error {
	networks, err := api.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(
			filters.Arg("name", NetworkName),
		),
	})

	if err != nil {
		return err
	}

	for _, network := range networks {
		err = api.NetworkRemove(ctx, network.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

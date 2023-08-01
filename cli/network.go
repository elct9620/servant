package cli

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const NetworkName = "servant"

func InstallNetwork(ctx context.Context, api client.APIClient) error {
	_, err := api.NetworkCreate(
		ctx,
		NetworkName,
		types.NetworkCreate{
			Driver:         "overlay",
			CheckDuplicate: true,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func UninstallNetwork(ctx context.Context, api client.APIClient) error {
	err := api.NetworkRemove(ctx, NetworkName)

	if err != nil {
		return err
	}

	return nil
}

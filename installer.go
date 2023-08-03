package servant

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"

	servantTypes "github.com/elct9620/servant/api/types"
)

type Installer struct {
}

func (i *Installer) Execute(ctx context.Context, api client.APIClient) error {
	fmt.Println("Installing the servant network...")
	network, err := i.InstallNetwork(ctx, api)
	if err != nil {
		return err
	}

	fmt.Println("Installing the servant controller...")
	_, err = i.InstallController(ctx, api, network.ID)
	if err != nil {
		return err
	}

	return nil
}

func (i *Installer) InstallNetwork(ctx context.Context, api client.NetworkAPIClient) (types.NetworkCreateResponse, error) {
	return api.NetworkCreate(
		ctx,
		NetworkName,
		types.NetworkCreate{
			Driver:         "overlay",
			CheckDuplicate: true,
			Labels: map[string]string{
				servantTypes.NameKey: NetworkName,
				servantTypes.TypeKey: servantTypes.TypeNetwork,
			},
		})
}

func (i *Installer) InstallController(ctx context.Context, api client.ServiceAPIClient, networkId string) (types.ServiceCreateResponse, error) {
	replicas := uint64(1)

	return api.ServiceCreate(
		ctx,
		swarm.ServiceSpec{
			Mode: swarm.ServiceMode{
				Replicated: &swarm.ReplicatedService{
					Replicas: &replicas,
				},
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image: fmt.Sprintf("%s:%s", ControllerImage, ControllerVersion),
					Mounts: []mount.Mount{
						{
							Type:   mount.TypeBind,
							Source: "/var/run/docker.sock",
							Target: "/var/run/docker.sock",
						},
					},
					Labels: map[string]string{
						servantTypes.NameKey: ControllerName,
						servantTypes.TypeKey: servantTypes.TypeController,
					},
				},
				Placement: &swarm.Placement{
					Constraints: []string{
						"node.role == manager",
					},
				},
				Networks: []swarm.NetworkAttachmentConfig{
					{
						Target: networkId,
					},
				},
			},
			Annotations: swarm.Annotations{
				Name: ControllerName,
				Labels: map[string]string{
					servantTypes.NameKey: ControllerName,
					servantTypes.TypeKey: servantTypes.TypeController,
				},
			},
		},
		types.ServiceCreateOptions{},
	)
}

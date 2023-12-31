package servant

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"

	servantTypes "github.com/elct9620/servant/api/types"
)

type Installer struct {
}

type InstallConfig struct {
	Version   string
	NetworkID string
}

func (i *Installer) Execute(ctx context.Context, api client.APIClient, config *InstallConfig) error {
	fmt.Println("Installing the servant network...")
	err := i.InstallNetwork(ctx, api, config)
	if err != nil {
		return err
	}

	fmt.Println("Installing the servant controller...")
	err = i.InstallController(ctx, api, config)
	if err != nil {
		return err
	}

	return nil
}

func (i *Installer) InstallNetwork(ctx context.Context, api client.NetworkAPIClient, config *InstallConfig) error {
	res, err := api.NetworkCreate(
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

	if errdefs.IsConflict(err) {
		res, err := api.NetworkInspect(ctx, NetworkName, types.NetworkInspectOptions{})
		if err != nil {
			return err
		}

		config.NetworkID = res.ID
		return nil
	}

	config.NetworkID = res.ID

	return err
}

func (i *Installer) InstallController(ctx context.Context, api client.ServiceAPIClient, config *InstallConfig) error {
	replicas := uint64(1)
	image := fmt.Sprintf("%s:%s", ControllerImage, ControllerVersion)

	if len(config.Version) > 0 {
		image = fmt.Sprintf("%s:%s", ControllerImage, config.Version)
	}

	_, err := api.ServiceCreate(
		ctx,
		swarm.ServiceSpec{
			Mode: swarm.ServiceMode{
				Replicated: &swarm.ReplicatedService{
					Replicas: &replicas,
				},
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image: image,
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
						Target: config.NetworkID,
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

	return err
}

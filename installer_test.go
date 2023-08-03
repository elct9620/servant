package servant_test

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type installFeature struct {
	client client.APIClient
}

func (i *installFeature) executeInstaller() error {
	return nil
}

func (i *installFeature) SetupScenario(ctx *godog.ScenarioContext) {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	i.client = dockerCli

	ctx.Step(`^I run "docker servant install$"`, i.executeInstaller)
	ctx.Step(`^the network should contains "([^"]*)$`, i.theNetworkShouldContains)
	ctx.Step(`^the service should contains "([^"]*)$`, i.theServiceShouldContains)
}

func (i *installFeature) theNetworkShouldContains(networkName string) error {
	_, err := i.client.NetworkInspect(context.Background(), networkName, types.NetworkInspectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (i *installFeature) theServiceShouldContains(serviceName string) error {
	services, err := i.client.ServiceList(context.Background(), types.ServiceListOptions{
		Filters: filters.NewArgs(filters.Arg("name", serviceName)),
	})
	if err != nil {
		return err
	}

	isFound := len(services) > 0
	if !isFound {
		return fmt.Errorf("service %s not found", serviceName)
	}

	return nil
}

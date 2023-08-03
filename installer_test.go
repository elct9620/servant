package servant_test

import (
	"context"
	"fmt"
	"time"

	"github.com/cucumber/godog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/docker/errdefs"
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

	ctx.Step("^I run `docker servant install`$", i.executeInstaller)
	ctx.Step(`^the network should contains "([^"]*)"$`, i.theNetworkShouldContains)
	ctx.Step(`^the service should contains "([^"]*)"$`, i.theServiceShouldContains)
}

func (i *installFeature) theNetworkShouldContains(ctx context.Context, networkName string) error {
	stepCtx, cancel := context.WithTimeout(ctx, stepWaitDuration)
	defer cancel()

	for {
		_, err := i.client.NetworkInspect(stepCtx, networkName, types.NetworkInspectOptions{})
		if errdefs.IsNotFound(err) {
			time.Sleep(1 * time.Second)
			continue
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func (i *installFeature) theServiceShouldContains(ctx context.Context, serviceName string) error {
	stepCtx, cancel := context.WithTimeout(ctx, stepWaitDuration)
	defer cancel()

	for {
		services, err := i.client.ServiceList(stepCtx, types.ServiceListOptions{
			Filters: filters.NewArgs(filters.Arg("name", serviceName)),
		})
		if errdefs.IsNotFound(err) {
			time.Sleep(1 * time.Second)
			continue
		}

		if err != nil {
			return err
		}

		isFound := len(services) > 0
		if !isFound {
			return fmt.Errorf("service %s not found", serviceName)
		}

		return nil
	}
}

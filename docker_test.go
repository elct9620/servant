package servant_test

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/google/go-cmp/cmp"
)

type dockerSteps struct {
	client client.APIClient
}

func (s *dockerSteps) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a docker client$`, s.aDockerClient)
	ctx.Step(`^I can see the following networks:$`, s.iCanSeeTheFollowingNetworks)
	ctx.Step(`^I can see the following services:$`, s.iCanSeeTheFollowingServices)
}

func (s *dockerSteps) aDockerClient() (err error) {
	s.client, err = client.NewClientWithOpts(client.FromEnv)
	return err
}

func (s *dockerSteps) iCanSeeTheFollowingNetworks(ctx context.Context, networks *godog.Table) error {
	expectedNetworks := []string{}
	for _, row := range networks.Rows[1:] {
		expectedNetworks = append(expectedNetworks, row.Cells[0].Value)
	}

	return asyncAssert(ctx, stepWaitDuration, func(ctx context.Context) error {
		networkList, err := s.client.NetworkList(ctx, types.NetworkListOptions{})
		if err != nil {
			return err
		}

		foundNetworkName := []string{}
		for _, network := range networkList {
			for _, name := range expectedNetworks {
				if network.Name == name {
					foundNetworkName = append(foundNetworkName, name)
				}
			}
		}

		if !cmp.Equal(expectedNetworks, foundNetworkName) {
			return fmt.Errorf("expected networks not found: %s", cmp.Diff(expectedNetworks, foundNetworkName))
		}

		return nil
	})
}

func (s *dockerSteps) iCanSeeTheFollowingServices(ctx context.Context, services *godog.Table) error {
	expectedServices := []string{}
	for _, row := range services.Rows[1:] {
		expectedServices = append(expectedServices, row.Cells[0].Value)
	}

	return asyncAssert(ctx, stepWaitDuration, func(ctx context.Context) error {
		serviceList, err := s.client.ServiceList(ctx, types.ServiceListOptions{})
		if err != nil {
			return err
		}

		foundServiceName := []string{}
		for _, service := range serviceList {
			for _, name := range expectedServices {
				if service.Spec.Name == name {
					foundServiceName = append(foundServiceName, name)
				}
			}
		}

		if !cmp.Equal(expectedServices, foundServiceName) {
			return fmt.Errorf("expected services not found: %s", cmp.Diff(expectedServices, foundServiceName))
		}

		return nil
	})
}

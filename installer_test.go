package servant_test

import (
	"context"
	"errors"

	"github.com/cucumber/godog"
	"github.com/docker/docker/client"
	"github.com/elct9620/servant"
)

type installSteps struct {
}

func (i *installSteps) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, sc *godog.Scenario, scErr error) (context.Context, error) {
		uninstaller := servant.Uninstaller{}

		dockerCli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return ctx, errors.Join(scErr, err)
		}

		err = uninstaller.Execute(ctx, dockerCli)
		if err != nil {
			return ctx, errors.Join(scErr, err)
		}

		return ctx, nil
	})
	ctx.Step(`^I run the installer$`, i.iRunTheInstaller)
}

func (i *installSteps) iRunTheInstaller(ctx context.Context) error {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	installer := &servant.Installer{}
	return installer.Execute(ctx, dockerCli, &servant.InstallConfig{})
}

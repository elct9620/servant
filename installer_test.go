package servant_test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/docker/docker/client"
	"github.com/elct9620/servant"
)

type installSteps struct {
}

func (i *installSteps) InitializeScenario(ctx *godog.ScenarioContext) {
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

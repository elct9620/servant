package servant_test

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
)

var e2eOptions = godog.Options{
	Format: "pretty",
	Tags:   "~@wip && ~@docker",
}

func init() {
	godog.BindCommandLineFlags("godog.", &e2eOptions)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := newApiFeature()
	api.SetupScenario(ctx)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	e2eOptions.Paths = []string{"features"}

	status := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &e2eOptions,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

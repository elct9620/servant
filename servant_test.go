package servant_test

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
)

const suiteSuccessCode = 0

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := newApiFeature()
	api.SetupScenario(ctx)

	install := &installFeature{}
	install.SetupScenario(ctx)
}

func TestFeatures(t *testing.T) {
	options := godog.Options{
		Format:   "pretty",
		Paths:    []string{"features"},
		Tags:     "~@wip && ~@docker",
		TestingT: t,
	}

	tags := os.Getenv("GODOG_TAGS")
	if len(tags) > 0 {
		options.Tags = tags
	}

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &options,
	}

	if st := suite.Run(); st != suiteSuccessCode {
		t.Errorf("Test suite failed with status code: %d", st)
	}
}

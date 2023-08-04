package servant_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/docker/docker/errdefs"
)

const suiteSuccessCode = 0
const stepWaitDuration = 30 * time.Second

type testStep interface {
	InitializeScenario(ctx *godog.ScenarioContext)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	steps := []testStep{
		newApiSteps(),
		&installSteps{},
		&dockerSteps{},
	}

	for _, step := range steps {
		step.InitializeScenario(ctx)
	}
}

func TestFeatures(t *testing.T) {
	isIntegration := os.Getenv("INTEGRATION") == "yes"
	if !isIntegration {
		t.SkipNow()
	}

	options := godog.Options{
		Format:   "pretty",
		Paths:    []string{"features"},
		Tags:     "~@wip",
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

type AsyncAssertFunc func(ctx context.Context) error

func asyncAssert(ctx context.Context, duration time.Duration, f AsyncAssertFunc) error {
	stepCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	for {
		err := f(stepCtx)
		if errdefs.IsNotFound(err) {
			time.Sleep(1 * time.Second)
			continue
		}

		if err != nil {
			return err
		}

		return stepCtx.Err()
	}
}

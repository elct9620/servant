package servant_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/elct9620/servant"
	"github.com/google/go-cmp/cmp"
)

type apiFeature struct {
	http.Handler
	resp *httptest.ResponseRecorder
}

func newApiFeature() *apiFeature {
	handler := servant.NewApi()

	return &apiFeature{
		Handler: handler,
	}
}

func (api *apiFeature) SetupScenario(ctx *godog.ScenarioContext) {
	api.resp = httptest.NewRecorder()

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resp = httptest.NewRecorder()
		return ctx, nil
	})

	ctx.Step(`^I make a GET request to "([^"]*)"$`, api.makeAGetRequestTo)
	ctx.Step(`^the response status code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response body should match JSON:$`, api.theResponseShouldMatchJSON)
}

func (api *apiFeature) makeAGetRequestTo(endpoint string) error {
	req := httptest.NewRequest("GET", endpoint, nil)
	api.ServeHTTP(api.resp, req)
	return nil
}

func (api *apiFeature) theResponseCodeShouldBe(code int) error {
	if api.resp.Code != code {
		return fmt.Errorf("the response code should be %d, but got %d", code, api.resp.Code)
	}
	return nil
}

func (api *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) error {
	var expected, actual interface{}
	if err := json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return err
	}

	if err := json.Unmarshal(api.resp.Body.Bytes(), &actual); err != nil {
		return err
	}

	if !cmp.Equal(expected, actual) {
		return fmt.Errorf("the response body should match the expected JSON, but got %s", cmp.Diff(expected, actual))
	}

	return nil
}

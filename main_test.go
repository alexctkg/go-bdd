package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-bdd/api"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/messages-go/v10"

	"github.com/cucumber/godog"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(*messages.Pickle) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendrequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch endpoint {
	case "/max-speed-allowed":
		api.GetMaxSpeedAllowed(a.resp, req)
	case "/{id}/last-speed":
		api.GetLastSpeed(a.resp, req)
	default:
		err = fmt.Errorf("unknown endpoint: %s", endpoint)
	}
	return
}
func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	var expected, actual []byte
	var data interface{}
	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return
	}
	if expected, err = json.Marshal(data); err != nil {
		return
	}
	actual = a.resp.Body.Bytes()
	if !bytes.Equal(actual, expected) {
		err = fmt.Errorf("expected json, does not match actual: %s", string(actual))
	}
	return
}

type serviceFeature struct {
}

func theResponseMessageShouldBeInsufficientDataPoints() error {
	return godog.ErrPending
}

func theResponseShouldContainAbruptBrakeEvent(arg1 int) error {
	return godog.ErrPending
}

func theResponseShouldContainBrakeEvent(arg1 int) error {
	return godog.ErrPending
}

func theResponseShouldContainNoEvent() error {
	return godog.ErrPending
}

func thereAreDataPoints(arg1 *messages.PickleStepArgument_PickleTable) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}

	s.BeforeScenario(api.resetResponse)

	s.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, api.iSendrequestTo)
	s.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`, api.theResponseShouldMatchJSON)
	//service
	// s.Step(`^I process points$`, service.iProcessPoints)
	// s.Step(`^the response message should be \'Insufficient data points\'$`, theResponseMessageShouldBeInsufficientDataPoints)
	// s.Step(`^the response should contain (\d+) abrupt brake event$`, theResponseShouldContainAbruptBrakeEvent)
	// s.Step(`^the response should contain (\d+) brake event$`, theResponseShouldContainBrakeEvent)
	// s.Step(`^the response should contain no event$`, theResponseShouldContainNoEvent)
	// s.Step(`^There are data points:$`, thereAreDataPoints)

}

package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ecoshub/taste/utils"
)

// hasOnlyRunMe to run only one case.
// loop all cases and find which one set 'OnlyRunThis' flag true
func (tt *Tester) hasOnlyRunMe(scenario []*Case) (*Case, bool) {
	for _, c := range scenario {
		if c.OnlyRunThis {
			return c, true
		}
	}
	return nil, false
}

func validateRequestStruct(c *Case) error {
	if c.Request == nil {
		return fmt.Errorf("error at test case '%s'. Field required 'Request'", c.Name)
	}
	if c.Request.Method == "" {
		return fmt.Errorf("error at test case '%s'. Field required 'Request.Method'", c.Name)
	}
	if c.Request.Path == "" {
		return fmt.Errorf("error at test case '%s'. Field required 'Request.Path'", c.Name)
	}
	return nil
}

func validateExpectStruct(c *Case) error {
	if c.Expect == nil {
		return fmt.Errorf("error at test case '%s'. Field required 'Expect'", c.Name)
	}
	return nil
}

func testTheCase(tester *Tester, c *Case, t *testing.T) {

	err := validateRequestStruct(c)
	if err != nil {
		t.Fatal(err)
	}

	err = validateExpectStruct(c)
	if err != nil {
		t.Fatal(err)
	}

	// if runBefore function defined run that before test
	if c.RunBefore != nil {
		c.RunBefore(t)
	}

	// if RunAfter function defined defer that to run after
	defer func() {
		if c.RunAfter != nil {
			c.RunAfter(t)
		}
	}()

	requestBody := []byte(c.Request.BodyString)

	// put in place stored value to request body
	requestBody, err = utils.InPlaceStoredValues(tester.store, requestBody)

	// check if any error ocurred during process
	utils.CheckExpectError(t, "request-body-process", err, nil)

	// create request with given request params
	req, err := http.NewRequest(c.Request.Method, c.Request.Path, bytes.NewBuffer(requestBody))
	utils.CheckExpectError(t, "request-creation", err, nil)

	// copy header of request and store it in to
	// http request header
	req.Header = c.Request.Header

	// do the request to tester handler
	resp, err := utils.Do(tester.handler, tester.ip, req)
	// check if any error ocurred during process
	utils.CheckExpectError(t, "handler-do", err, nil)

	// local failed flag
	failed := false

	// check if response status code is expected
	statusFailed := utils.CheckEqualLog(t, "response-status-code", resp.StatusCode, c.Expect.Status)

	// if not expected set local failed flag to false
	failed = failed || statusFailed

	if c.CheckHeader {
		// check if response header code is expected
		headerFailed := utils.CheckEqualLog(t, "response-header", resp.Header, c.Expect.Header)
		// if not expected set local failed flag to false
		failed = failed || headerFailed
	}

	// read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("response body read error. err:", err)
	}
	defer resp.Body.Close()

	expectedBody := []byte(c.Expect.BodyString)

	// put in place stored value to expected request body
	expectedBody, err = utils.InPlaceStoredValues(tester.store, expectedBody)
	// check if any error ocurred during process
	utils.CheckExpectError(t, "expect-body-process", err, nil)

	// validate responseBody with expected body
	err = utils.Validate(expectedBody, responseBody)
	if err == nil && c.Expect.Error != nil {
		utils.Fail(t, "error-expected", err, c.Expect.Error)
	}
	if err != nil && c.Expect.Error == nil {
		t.Fatal(" 'response-body-validation'", err)
	}
	if err != nil && c.Expect.Error != nil {
		utils.CheckExpectError(t, "error", err, c.Expect.Error)
	}

	// if err != nil {
	// 	if c.Expect.Error == nil {
	// 		utils.Fail(t, "not-error-expected", err, c.Expect.Error)
	// 		return
	// 	}
	// 	t.Fatal(" 'response-body-validation'", err)
	// } else {
	// 	if c.Expect.Error != nil {
	// 		return
	// 	}
	// }

	// if any check failed above fail the test
	if failed {
		t.Fail()
	}

	if c.StoreResponse {
		tester.store[c.Name] = requestBody
	}

}

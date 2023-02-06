package taste

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func tasteIt(tt *ServerTester, c *HTTPTestCase) {

	// validate given request struct
	err := validateRequestStruct(c)
	if err != nil {
		tt.t.Fatal(err)
	}

	// validate given expect struct
	err = validateExpectStruct(c)
	if err != nil {
		tt.t.Fatal(err)
	}

	// if runBefore function defined run that before test
	if c.RunBefore != nil {
		c.RunBefore(tt.t)
	}

	// if RunAfter function defined defer that to run after
	defer func() {
		if c.RunAfter != nil {
			c.RunAfter(tt.t)
		}
	}()

	requestBody := []byte(c.Request.Body)

	// put in place stored value to request body
	requestBody, err = inPlaceStoredValues(tt.store, requestBody)

	// check if any error ocurred during process
	checkExpectError(tt.t, "request-body-process", err, nil)

	// create request with given request params
	req, err := http.NewRequest(c.Request.Method, c.Request.Path, bytes.NewBuffer(requestBody))
	checkExpectError(tt.t, "request-creation", err, nil)

	// copy header of request and store it in to
	// http request header
	req.Header = c.Request.Header

	// do the request to tester handler
	resp, err := do(tt.handler, tt.ip, req)
	// check if any error ocurred during process
	checkExpectError(tt.t, "handler-do", err, nil)

	// local failed flag
	failed := false

	// check if response status code is expected
	statusFailed := checkEqualLog(tt.t, "response-status-code", resp.StatusCode, c.Expect.Status)

	// if not expected set local failed flag to false
	failed = failed || statusFailed

	if c.CheckHeader {
		// check if response header code is expected
		headerFailed := checkEqualLog(tt.t, "response-header", resp.Header, c.Expect.Header)
		// if not expected set local failed flag to false
		failed = failed || headerFailed
	}

	// read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		tt.t.Fatal("response body read error. err:", err)
	}
	defer resp.Body.Close()

	expectedBody := []byte(c.Expect.Body)

	// put in place stored value to expected request body
	expectedBody, err = inPlaceStoredValues(tt.store, expectedBody)
	// check if any error ocurred during process
	checkExpectError(tt.t, "expect-body-process", err, nil)

	// validate responseBody with expected body
	err = validate(expectedBody, responseBody)
	if err == nil && c.Expect.Error != nil {
		fail(tt.t, "error-expected", err, c.Expect.Error)
	}
	if err != nil && c.Expect.Error == nil {
		tt.t.Fatal(" 'response-body-validation'", err)
	}
	if err != nil && c.Expect.Error != nil {
		checkExpectError(tt.t, "error", err, c.Expect.Error)
	}

	// if any check failed above fail the test
	if failed {
		tt.t.Fail()
	}

	if c.StoreResponse {
		tt.store[c.Name] = requestBody
	}

}

// hasOnlyRunMe to run only one case.
// loop all cases and find which one set 'OnlyRunThis' flag true
func (tt *ServerTester) hasOnlyRunMe(scenario []*HTTPTestCase) int {
	for i, c := range scenario {
		if c.OnlyRunThis {
			return i
		}
	}
	return -1
}

// validateRequestStruct field control for Request struct
func validateRequestStruct(c *HTTPTestCase) error {
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

// validateExpectStruct field control for Expect struct
func validateExpectStruct(c *HTTPTestCase) error {
	if c.Expect == nil {
		return fmt.Errorf("error at test case '%s'. Field required 'Expect'", c.Name)
	}
	return nil
}

package server

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/ecoshub/taste/utils"
)

type Tester struct {
	handler http.Handler
	store   map[string][]byte
}

// NewTester creates a new instance of Tester with the provided http.Handler and an empty store.
func NewTester(handler http.Handler) *Tester {
	return &Tester{
		handler: handler,
		store:   make(map[string][]byte),
	}
}

// runCase runs a single test case and checks the response against the expected values.
func (tt *Tester) runCase(t *testing.T, c *Case) {
	// Run the "RunBefore" function if it exists.
	if c.RunBefore != nil {
		c.RunBefore(t, c)
	}

	// Defer the "RunAfter" function if it exists.
	defer func() {
		if c.RunAfter != nil {
			c.RunAfter(t, c)
		}
	}()

	// Process the request body using the store.
	processedRequestBody, err := utils.ProcessBody(tt.store, c.Request.Body)
	utils.CheckExpectError(t, "request-body-process", err, nil)

	// Create a new http.Request with the processed request body.
	req, err := http.NewRequest(c.Request.Method, c.Request.RequestURI, bytes.NewBuffer(processedRequestBody))
	utils.CheckExpectError(t, "request-creation", err, nil)

	// Set the request headers.
	req.Header = c.Request.Header

	// Send the request and get the response.
	resp := utils.Do(tt.handler, req)

	// Check if the response status code matches the expected value.
	utils.CheckEqual(t, "response-status-code", resp.StatusCode, c.Response.Status)

	// Check if the response headers match the expected values.
	if len(c.Response.Header) > 0 {
		utils.CheckEqual(t, "response-header", resp.Header, c.Response.Header)
	}

	// Read the response body.
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("response body read error. err:", err)
	}
	defer resp.Body.Close()

	// Process the response body using the store.
	expectedBody, err := utils.ProcessBody(tt.store, c.Response.Body)
	utils.CheckExpectError(t, "expect-body-process", err, nil)

	// Check if the response body matches the expected value.
	if len(c.Response.Body) == 0 && len(responseBody) != 0 {
		t.Fatalf("expected nothing but got something. body: %s", responseBody)
	}

	// Validate the response body against the expected body.
	err = utils.Validate(expectedBody, responseBody)
	// If there is an error...
	if err != nil {
		// Check if the error is expected.
		if c.Response.Error != nil {
			utils.CheckEqual(t, "error", err, c.Response.Error)
			return
		}
		// Otherwise, fail the test.
		t.Fatalf("err: %v. expected: %s, got: %s", err, expectedBody, responseBody)
	}

	// If the "StoreResponse" flag is set, store the response body in the store.
	if c.StoreResponse {
		tt.store[c.Name] = responseBody
	}
}

// Run runs the provided scenario using the attached http.Handler.
func (tt *Tester) Run(t *testing.T, scenario []*Case) {
	// Check if there is a handler attached to the tester.
	if tt.handler == nil {
		t.Fatal("there is no handler to test this scenario. please attach a handler with 'AttachHandler' function")
	}

	// Check if there is a test case with the "OnlyRunThis" flag set.
	for _, c := range scenario {
		if c.OnlyRunThis {
			t.Logf("RUN [ONLY]\t%s\n", c.Name)
			t.Run(c.Name, func(t *testing.T) {
				tt.runCase(t, c)
			})
			return
		}
	}

	// Otherwise, run all test cases in the provided order.
	for _, c := range scenario {
		t.Run(c.Name, func(t *testing.T) {
			tt.runCase(t, c)
		})
	}
}

// ResetStore resets the store to an empty map.
func (tt *Tester) ResetStore() {
	tt.store = make(map[string][]byte)
}

// StoreKeyValue stores a key-value pair in the store.
func (tt *Tester) StoreKeyValue(key string, body []byte) {
	tt.store[key] = body
}

// StoreKeyValueString stores a key-value pair in the store, with the value provided as a string.
func (tt *Tester) StoreKeyValueString(key, value string) {
	tt.StoreKeyValue(key, []byte(value))
}

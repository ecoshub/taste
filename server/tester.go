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

func NewTester(handler http.Handler) *Tester {
	return &Tester{
		handler: handler,
		store:   make(map[string][]byte)}
}

func (tt *Tester) ResetStore() {
	tt.store = make(map[string][]byte)
}

func (tt *Tester) StoreKeyValue(key string, body []byte) {
	tt.store[key] = body
}

func (tt *Tester) StoreKeyValueString(key, value string) {
	tt.StoreKeyValue(key, []byte(value))
}

func (tt *Tester) Run(t *testing.T, scenario []*Case) {
	if tt.handler == nil {
		t.Fatal("there is  no handler to test this scenario. please attach a handler with 'AttachHandler' function")
	}

	for _, c := range scenario {
		if c.OnlyRunThis {
			t.Logf("RUN [ONLY]\t%s\n", c.Name)
			t.Run(c.Name, func(t *testing.T) {
				tt.runCase(t, c)
			})
			return
		}
	}

	for _, c := range scenario {
		t.Run(c.Name, func(t *testing.T) {
			tt.runCase(t, c)
		})
	}
}

func (tt *Tester) runCase(t *testing.T, c *Case) {
	if c.RunBefore != nil {
		c.RunBefore(t, c)
	}

	defer func() {
		if c.RunAfter != nil {
			c.RunAfter(t, c)
		}
	}()

	processedRequestBody, err := utils.ProcessBody(tt.store, c.Request.Body)
	utils.CheckExpectError(t, "request-body-process", err, nil)

	req, err := http.NewRequest(c.Request.Method, c.Request.RequestURI, bytes.NewBuffer(processedRequestBody))
	utils.CheckExpectError(t, "request-creation", err, nil)

	req.Header = c.Request.Header

	resp := utils.Do(tt.handler, req)

	utils.CheckEqual(t, "response-status-code", resp.StatusCode, c.Response.Status)

	if len(c.Response.Header) > 0 {
		utils.CheckEqual(t, "response-header", resp.Header, c.Response.Header)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("response body read error. err:", err)
	}
	defer resp.Body.Close()

	expectedBody, err := utils.ProcessBody(tt.store, responseBody)
	utils.CheckExpectError(t, "expect-body-process", err, nil)

	if len(c.Response.Body) == 0 && len(responseBody) != 0 {
		t.Fatalf("expected nothing but got something. body: %s", responseBody)
	}

	err = utils.Validate(expectedBody, responseBody)
	// got error
	if err != nil {
		// expecting error
		if c.Response.Error != nil {
			utils.CheckEqual(t, "error", err, c.Response.Error)
			return
		}
		t.Fatalf("err: %v. expected: %s, got: %s", err, expectedBody, responseBody)
	}

	if c.StoreResponse {
		tt.store[c.Name] = responseBody
	}
}

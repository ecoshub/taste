package server

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/ecoshub/taste/utils"
)

func (sc *Tester) hasOnlyRunMe() (*Case, bool) {
	for _, c := range sc.scenario {
		if c.OnlyRunThis {
			return c, true
		}
	}
	return nil, false
}

func run(tester *Tester, c *Case, t *testing.T) {
	if c.RunBefore != nil {
		c.RunBefore(t)
	}

	defer func() {
		if c.RunAfter != nil {
			c.RunAfter(t)
		}
	}()

	body := resolveBody(c.Request.Body, c.Request.BodyString)

	var err error
	body, err = utils.ProcessBody(tester.store, body)
	utils.CheckExpectError(t, "request-body-process", err, nil)

	buff := bytes.NewBuffer(body)

	req, err := http.NewRequest(c.Request.Method, c.Request.Path, buff)
	utils.CheckExpectError(t, "request-creation", err, nil)

	req.Header = c.Request.Header

	resp, err := utils.Do(tester.handler, tester.ip, req)
	utils.CheckExpectError(t, "handler-do", err, nil)

	utils.CheckEqual(t, "response-status-code", resp.StatusCode, c.Expect.Status)

	if len(c.Expect.Header) > 0 {
		utils.CheckEqual(t, "response-header", resp.Header, c.Expect.Header)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("response body read error. err:", err)
	}
	defer resp.Body.Close()

	expectedBody := resolveBody(c.Expect.Body, c.Expect.BodyString)

	expectedBody, err = utils.ProcessBody(tester.store, expectedBody)
	utils.CheckExpectError(t, "expect-body-process", err, nil)

	err = utils.Validate(expectedBody, body)
	// got error
	if err != nil {
		// expecting error
		if c.Expect.Error != nil {
			utils.CheckEqual(t, "error", err, c.Expect.Error)
			return
		}
		t.Fatalf("err: %v. expected: %s, got: %s", err, expectedBody, body)
	}

	if c.StoreResponse {
		tester.store[c.Name] = body
	}

}

func resolveBody(body []byte, bodyString string) []byte {
	if body == nil {
		if bodyString != "" {
			return []byte(bodyString)
		} else {
			return []byte{}
		}
	} else {
		return body
	}
}

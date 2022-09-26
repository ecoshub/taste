package taste

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/ecoshub/taste/utils"
)

func (sc *Scenario) hasOnlyRunMe() (*Case, bool) {
	for _, c := range sc.cases {
		if c.OnlyRunMe {
			return c, true
		}
	}
	return nil, false
}

func run(sc *Scenario, c *Case, t *testing.T) {
	if sc.runBefore != nil {
		sc.runBefore(t)
	}

	defer func() {
		if sc.runAfter != nil {
			sc.runAfter(t)
		}
	}()

	buff := resolveBody(c.Request)

	req, err := http.NewRequest(c.Request.Method, c.Request.URL, buff)
	utils.CheckExpectError(t, "request-creation", err, nil)

	req.Header = c.Request.Header

	resp, err := utils.Do(sc.server, sc.mockIP, req)
	utils.CheckExpectError(t, "handler-do", err, nil)

	utils.CheckEqual(t, "response-status-code", resp.StatusCode, c.Expect.Status)

	if len(c.Expect.Header) > 0 {
		utils.CheckEqual(t, "response-header", resp.Header, c.Expect.Header)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("response body read error. err:", err)
	}
	defer resp.Body.Close()

	if c.Expect.Body == nil {
		if c.Expect.BodyString != "" {
			utils.CheckEqual(t, "response-body", string(body), c.Expect.BodyString)
		}
	} else {
		utils.CheckEqual(t, "response-body", string(body), string(c.Expect.Body))
	}

}

func resolveBody(r *Request) *bytes.Buffer {
	if r.Body == nil {
		if r.BodyString != "" {
			return bytes.NewBufferString(r.BodyString)
		} else {
			return &bytes.Buffer{}
		}
	} else {
		return bytes.NewBuffer(r.Body)
	}
}

package taste

import (
	"net/http"
	"testing"
)

var (
	// HeaderContentPlainText standard plain text header
	HeaderContentPlainText = http.Header{"Content-Type": []string{"text/plain"}}

	// HeaderContentPlainText standard application json header
	HeaderContentApplicationJSON = http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
)

// HTTPTestCase test case for http server test
type HTTPTestCase struct {
	// Name name of test case
	Name string

	// Request request is request scheme to test.
	// this field is required
	Request *Request

	// Expect expect is expected response scheme
	// this field is required
	Expect *Expect

	// OnlyRunThis if this flag enabled, test system will run only this
	// if multiple test case enabled this flag, system will run the first occurred test case only
	OnlyRunThis bool

	// StoreResponse store response body to tester for later use
	// to use stored values use double angle notation
	// see documentation for more information
	StoreResponse bool

	// CheckHeader if this enabled test system will compare expected headers with response headers
	CheckHeader bool

	// RunBefore this function will run before testing if its defined
	RunBefore func(t *testing.T)

	// RunBefore this function will run after testing if its defined
	RunAfter func(t *testing.T)
}

// Request test case request scheme
type Request struct {
	// Method http request method
	Method string
	// Header http headers
	Header http.Header
	// Path http request path
	Path string
	// Body request body
	Body string
}

// Request test case expectation scheme
type Expect struct {
	// Status expected response status
	Status int
	// Body expected response body
	Body string
	// Header expected response headers
	Header http.Header
	// Error expected error
	Error error
}

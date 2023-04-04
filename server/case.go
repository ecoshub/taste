package server

import (
	"testing"
)

// Case represents an individual test case
type Case struct {
	// Name identifies the name of the test case
	Name string
	// OnlyRunThis indicates whether this test case should be the only one run
	OnlyRunThis bool
	// StoreResponse indicates whether the response should be stored for later use
	StoreResponse bool
	// Request defines the HTTP request to be made
	Request *Request
	// Response defines the expected HTTP response
	Response *Response
	// RunBefore is a function to run before the test case
	RunBefore func(t *testing.T, c *Case)
	// RunAfter is a function to run after the test case
	RunAfter func(t *testing.T, c *Case)
}

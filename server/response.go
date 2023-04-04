package server

import "net/http"

// Response represents an expected HTTP response
type Response struct {
	// Status is the expected HTTP status code (e.g., 200, 404)
	Status int
	// Body is the expected response body as a byte slice
	Body []byte
	// Header is the expected HTTP headers in the response
	Header http.Header
	// Error is an expected error, if any (e.g., an error parsing the response body)
	Error error
}

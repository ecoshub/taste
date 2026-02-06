package server

import (
	"net/http"
	"net/url"
)

// Request represents an HTTP request
type Request struct {
	// Method is the HTTP method to use (e.g., GET, POST)
	Method string
	// RequestURI is the URI to request (e.g., /path/to/resource)
	RequestURI string
	// Header is the HTTP headers to include with the request
	Header http.Header
	// Query is the query parameters to include with the request
	Query url.Values
	// Body is the request body as a byte slice
	Body []byte
	// BodyString is the request body as a string
	BodyString string
}

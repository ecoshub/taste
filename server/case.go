package server

import (
	"net/http"
	"net/url"
	"testing"
)

type scenario []*Case

type Case struct {
	Name        string
	OnlyRunThis bool
	Request     *Request
	Expect      *Expect
	RunBefore   func(t *testing.T)
	RunAfter    func(t *testing.T)
}

type Request struct {
	Method     string
	Path       string
	Header     http.Header
	Query      url.Values
	Body       []byte
	BodyString string
}

type Expect struct {
	Status     int
	Body       []byte
	BodyString string
	Header     http.Header
	Error      error
}

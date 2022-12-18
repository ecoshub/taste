package server

import (
	"net/http"
	"testing"
)

var (
	HeaderContentPlainText       = http.Header{"Content-Type": []string{"text/plain"}}
	HeaderContentApplicationJSON = http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
)

type Case struct {
	Name          string
	OnlyRunThis   bool
	StoreResponse bool
	CheckHeader   bool
	Request       *Request
	Expect        *Expect
	RunBefore     func(t *testing.T)
	RunAfter      func(t *testing.T)
}

type Request struct {
	Method     string
	Path       string
	Header     http.Header
	BodyString string
}

type Expect struct {
	Status     int
	Body       []byte
	BodyString string
	Header     http.Header
	Error      error
}

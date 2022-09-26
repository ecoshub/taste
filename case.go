package taste

import (
	"net/http"
	"net/url"
)

type Case struct {
	Name      string
	Request   *Request
	Expect    *Expect
	OnlyRunMe bool
}

type Request struct {
	Method     string
	URL        string
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
}

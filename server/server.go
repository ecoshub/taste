package server

import (
	"net/http"
)

const (
	defaultMockIP string = "127.0.0.1"
)

type HTTPMock struct {
	mux *http.ServeMux
	ip  string
}

func NewHTTPServer(mockIPoptional ...string) *HTTPMock {
	mockIP := defaultMockIP
	if len(mockIPoptional) == 1 {
		mockIP = mockIPoptional[0]
	}
	mux := http.NewServeMux()
	return &HTTPMock{mux: mux, ip: mockIP}
}

func (s *HTTPMock) Handler() http.Handler {
	return s.mux
}

func (s *HTTPMock) Handle(method, path string, handlerFunc http.HandlerFunc) {
	s.mux.HandleFunc(path, methodMiddleware(method, handlerFunc))
}

func methodMiddleware(method string, f http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		f(w, r)
	}
}

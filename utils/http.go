package utils

import (
	"net/http"
	"net/http/httptest"
)

// Do makes an HTTP request to the given handler and returns the response
func Do(handler http.Handler, req *http.Request) *http.Response {
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w.Result()
}

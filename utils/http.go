package utils

import (
	"net/http"
	"net/http/httptest"
)

func Do(handler http.Handler, mockIP string, req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	req.RemoteAddr = mockIP
	handler.ServeHTTP(w, req)
	return w.Result(), nil
}

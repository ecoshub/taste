package taste

import (
	"net/http"
	"net/http/httptest"
)

// do do request to given handler
func do(handler http.Handler, mockIP string, req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	req.RemoteAddr = mockIP
	handler.ServeHTTP(w, req)
	return w.Result(), nil
}

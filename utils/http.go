package utils

import (
	"net/http"
	"net/http/httptest"
)

func HeaderPair(keyValuePairs ...string) http.Header {
	if len(keyValuePairs)%2 != 0 {
		return http.Header{"error": []string{"key value pair count is not even"}}
	}
	h := http.Header{}
	for i := 0; i < len(keyValuePairs); i += 2 {
		key := keyValuePairs[i]
		val := keyValuePairs[i+1]
		h.Add(key, val)
	}
	return h
}

// Do makes an HTTP request to the given handler and returns the response
func Do(handler http.Handler, req *http.Request) *http.Response {
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w.Result()
}

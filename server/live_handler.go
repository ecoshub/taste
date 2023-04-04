package server

import (
	"io"
	"net/http"
)

type LiveServerHandler struct{}

func (lh *LiveServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		if resp != nil {
			w.WriteHeader(resp.StatusCode)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	// copy headers
	for k, vals := range resp.Header {
		for _, v := range vals {
			w.Header().Add(k, v)
		}
	}

	io.Copy(w, resp.Body)
}

func NewLiveServerTester() *Tester {
	return &Tester{
		handler: &LiveServerHandler{},
		store:   make(map[string][]byte),
	}
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func ExampleServer() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/version", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "plain/text")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("v1.0.0"))
	})
	m.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		user := fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, username)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(user))
	})
	m.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})
	return m
}

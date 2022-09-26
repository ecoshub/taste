package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ecoshub/taste"
	"github.com/ecoshub/taste/example"
)

func exampleDefaultServer() *taste.HTTPMockServer {
	s := taste.NewHTTPServer("127.0.0.1")
	s.Handle("GET", "/api/v1/version", versionHandler)
	s.Handle("GET", "/api/v1/users", usersHandler)
	s.Handle("GET", "/api/v1/user", userHandler)
	s.Handle("POST", "/api/v1/user/new", newUserHandler)
	return s
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("v1.0.0"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(example.MarshalDiscardError(example.Users))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	user, exists := example.GetUser(name)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("404 page not found"))
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(user.Marshal())
}

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write(example.MarshalDiscardError(map[string]string{"error": err.Error()}))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user *example.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.Write(example.MarshalDiscardError(map[string]string{"error": err.Error()}))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	example.AddUser(user)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(example.MarshalDiscardError(example.Users))
}

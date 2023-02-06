package main_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ecoshub/taste"
)

var (
	Scenario = []*taste.HTTPTestCase{
		{
			StoreResponse: true,
			Name:          "get_random",
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/random",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   `{"random|string":"*"}`,
			},
		},
		{
			StoreResponse: true,
			Name:          "echo_test",
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/echo",
				Body:   `{"id":"<<test_1.random>>","name":"eco"}`,
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   `{"id":"<<test_1.random>>","name":"eco"}`,
			},
		},
		{
			StoreResponse: true,
			Name:          "test_3",
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/echo",
				Body:   `{"id":"<<test_1.random>>","name":"<<test_2.name>>"}`,
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   `{"id":"<<test_2.id>>","name":"<<test_2.name>>"}`,
			},
		},
	}
)

func NewServer() *taste.MockServer {
	s := taste.NewMockServer("127.0.0.1")
	s.Handle("GET", "/api/v1/echo", echoHandler)
	s.Handle("GET", "/api/v1/random", randomHandler)
	return s
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body := []byte(fmt.Sprintf(`{"random":"%s"}`, taste.RandomHash(16)))
	w.Write(body)
}

func TestCore(t *testing.T) {

	// create new server
	s := NewServer()

	// create a tester with server handler and scenario
	tester := taste.NewHTTPServerTester(t, s.Handler())

	// run the scenario
	tester.Run(Scenario)
}

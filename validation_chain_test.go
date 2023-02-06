package taste

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

var (
	Scenario = []*HTTPTestCase{
		{
			StoreResponse: true,
			Name:          "test_1",
			Request: &Request{
				Method: http.MethodGet,
				Path:   "/api/v1/random",
			},
			Expect: &Expect{
				Status: http.StatusOK,
				Body:   `{"random|string":"*"}`,
			},
		},
		{
			StoreResponse: true,
			Name:          "test_2",
			Request: &Request{
				Method: http.MethodGet,
				Path:   "/api/v1/echo",
				Body:   `{"id":"<<test_1.random>>","name":"eco"}`,
			},
			Expect: &Expect{
				Status: http.StatusOK,
				Body:   `{"id":"<<test_1.random>>","name":"eco"}`,
			},
		},
		{
			StoreResponse: true,
			Name:          "test_3",
			Request: &Request{
				Method: http.MethodGet,
				Path:   "/api/v1/echo",
				Body:   `{"id":"<<test_1.random>>","name":"<<test_2.name>>"}`,
			},
			Expect: &Expect{
				Status: http.StatusOK,
				Body:   `{"id":"<<test_2.id>>","name":"<<test_2.name>>"}`,
			},
		},
	}
)

func NewServer() *MockServer {
	s := NewMockServer("127.0.0.1")
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
	body := []byte(fmt.Sprintf(`{"random":"%s"}`, RandomHash(16)))
	w.Write(body)
}

func TestCore(t *testing.T) {

	// create new server
	s := NewServer()

	// create a tester with server handler and scenario
	tester := NewHTTPServerTester(t, s.Handler())

	// run the scenario
	tester.Run(Scenario)
}

package chain_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	example "github.com/ecoshub/taste/example/server"
	"github.com/ecoshub/taste/server"
	"github.com/ecoshub/taste/utils"
)

var (
	Scenario = []*server.Case{
		{
			StoreResponse: true,
			Name:          "test_1",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/random",
			},
			Expect: &server.Expect{
				Status:     http.StatusOK,
				BodyString: `{"random|string":"*"}`,
			},
		},
		{
			StoreResponse: true,
			Name:          "test_2",
			Request: &server.Request{
				Method:     http.MethodGet,
				Path:       "/api/v1/echo",
				BodyString: `{"id":"<<test_1.random>>","name":"eco"}`,
			},
			Expect: &server.Expect{
				Status:     http.StatusOK,
				BodyString: `{"id":"<<test_1.random>>","name":"eco"}`,
			},
		},
		{
			StoreResponse: true,
			Name:          "test_3",
			Request: &server.Request{
				Method:     http.MethodGet,
				Path:       "/api/v1/echo",
				BodyString: `{"id":"<<test_1.random>>","name":"<<test_2.name>>"}`,
			},
			Expect: &server.Expect{
				Status:     http.StatusOK,
				BodyString: `{"id":"<<test_2.id>>","name":"<<test_2.name>>"}`,
			},
		},
	}
)

func NewServer() *server.HTTPMock {
	s := server.NewHTTPServer("127.0.0.1")
	s.Handle("GET", "/api/v1/echo", echoHandler)
	s.Handle("GET", "/api/v1/random", randomHandler)
	return s
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write(example.MarshalDiscardError(map[string]string{"error": err.Error()}))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	random := utils.RandomHash(16)
	body := []byte(fmt.Sprintf(`{"random":"%s"}`, random))
	w.Write(body)
}

func TestCore(t *testing.T) {

	s := NewServer()

	// create a tester with server handler and scenario
	tester := server.NewTester(Scenario, s.Handler())

	// run the scenario
	tester.Run(t)
}

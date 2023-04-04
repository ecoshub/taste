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
	scenario = []*server.Case{
		{
			StoreResponse: true,
			Name:          "test_1",
			Request: &server.Request{
				Method:     http.MethodGet,
				RequestURI: "/api/v1/random",
			},
			Response: &server.Response{
				Status: http.StatusOK,
				Body:   []byte(`{"random|string":"*"}`),
			},
		},
		{
			StoreResponse: true,
			Name:          "test_2",
			Request: &server.Request{
				Method:     http.MethodGet,
				RequestURI: "/api/v1/echo",
				Body:       []byte(`{"id":"<<test_1.random>>","name":"eco"}`),
			},
			Response: &server.Response{
				Status: http.StatusOK,
				Body:   []byte(`{"id":"<<test_1.random>>","name":"eco"}`),
			},
		},
		{
			StoreResponse: true,
			Name:          "test_3",
			Request: &server.Request{
				Method:     http.MethodGet,
				RequestURI: "/api/v1/echo",
				Body:       []byte(`{"id":"<<test_1.random>>","name":"<<test_2.name>>"}`),
			},
			Response: &server.Response{
				Status: http.StatusOK,
				Body:   []byte(`{"id":"<<test_2.id>>","name":"<<test_2.name>>"}`),
			},
		},
	}
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/echo", echoHandler)
	mux.HandleFunc("/api/v1/random", randomHandler)
	return mux
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

	mux := NewServer()

	// create a tester with server handler and scenario
	tester := server.NewTester(mux)

	// run the scenario
	tester.Run(t, scenario)
}

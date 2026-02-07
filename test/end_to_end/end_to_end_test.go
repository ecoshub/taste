package end_to_end_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

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
				BodyString: `{"id":"<<test_1.random>>","name":"<<test_2.name>>"}`,
			},
			Response: &server.Response{
				Status:     http.StatusOK,
				BodyString: `{"id":"<<test_2.id>>","name":"<<test_2.name>>"}`,
			},
		},
		{
			StoreResponse: true,
			Name:          "test_4",
			Request: &server.Request{
				Method:     http.MethodGet,
				RequestURI: "/api/v1/echo",
				BodyString: `{"id":"61"}`,
			},
			Response: &server.Response{
				Status:     http.StatusOK,
				BodyString: `{"id":"61"}`,
			},
		},
		{
			Name: "test_5",
			Request: &server.Request{
				Method:     http.MethodGet,
				RequestURI: "/api/v1/null",
				BodyString: `{"id":"62"}`,
			},
			Response: &server.Response{
				Status:     http.StatusOK,
				BodyString: "*",
			},
		},
		{
			Name: "test_6",
			Request: &server.Request{
				Method:     http.MethodGet,
				RequestURI: "/api/v1/null",
				BodyString: `{"id":"61"}`,
			},
			Response: &server.Response{
				Status:     http.StatusOK,
				BodyString: `{}`,
			},
		},
		{
			Name: "test_7",
			Request: &server.Request{
				Method:     http.MethodGet,
				RequestURI: "/api/v1/echo",
				BodyString: `{"id":"61"}`,
				Header: http.Header{
					"hello": []string{"world"},
					"test":  []string{"<<test_4.id>>"},
				},
			},
			Response: &server.Response{
				Status: http.StatusOK,
				Header: utils.HeaderPair(
					"Content-Type", "text/plain; charset=utf-8",
					"hello", "world",
					"test", "61",
				),
				BodyString: `{"id":"61"}`,
			},
		},
	}
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/echo", echoHandler)
	mux.HandleFunc("/api/v1/random", randomHandler)
	mux.HandleFunc("/api/v1/null", nullHandler)
	return mux
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for key, values := range r.Header {
		for _, val := range values {
			w.Header().Add(key, val)
		}
	}
	w.Write(body)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	random := utils.RandomHash(16)
	body := []byte(fmt.Sprintf(`{"random":"%s"}`, random))
	w.Write(body)
}

func nullHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}`))
}

func Test(t *testing.T) {

	mux := NewServer()

	// create a tester with server handler and scenario
	tester := server.NewTester(mux)

	// run the scenario
	tester.Run(t, scenario)
}

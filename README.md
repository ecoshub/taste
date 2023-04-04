## Taste
*A taste of simplicity.*

taste is a simple (but powerful) table driven testing tool.

## Capabilities:

- HTTP Server Testing: 

It allows you to define a scenario of HTTP requests and expected responses, and will run the scenario and perform assertions to ensure that the server is behaving as expected.

- Unit testing:

 It offers to define and execute tests, as well as to validate expected outcomes. Its flexible and straightforward design makes writing and managing unit tests easier. It allows table-driven testing to execute multiple tests in a single function through a slice of structs.


## HTTP Testing utils

### JSON body Validation with scheme

- The validation is done against a JSON validation scheme, which specifies the expected structure and content of the data.

- The validation scheme uses a specific syntax to specify the structure and content of a JSON object. The syntax is `"field_name|type": "value"`, where field_name is the name of the field, type is the data type of the field, and value is the expected value of the field.

- The validation scheme supports wildcards, which can be used to match any value or data type. The wildcard is denoted by *.

- The validate function returns an error if the data does not match the scheme, and returns nil if the data matches the scheme.

### Double bracket notation

Double angle bracket notation (e.g. <<test_1.random>>) is used to reference the value of a field in a previous HTTP response.

This allows you to reuse values from previous responses in subsequent requests, which can be useful for testing scenarios where the server relies on data from previous requests.


## Example (HTTP Testing):

```go
package server_test

import (
	"fmt"
	"io/ioutil"
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
	mux.HandleFunc("/api/v1/random", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(fmt.Sprintf(`{"random":"%s"}`, utils.RandomHash(16))))
	})
	mux.HandleFunc("/api/v1/echo", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
			return
		}
		w.Write(body)
	})
	return mux
}

func TestCustomServer(t *testing.T) {
	// your server implementation 
	s := NewServer()

	// create a tester with server handler and scenario
	// pass your server handler
	// also you can pass your "*gin.Engine" as handler
	tester := server.NewTester(s)

	// run the scenario
	tester.Run(t, scenario)
}
```

## Example (Unit testing):

```go
package unit

import (
	"errors"
	"testing"

	"github.com/ecoshub/taste/unit"
)

var (
	scenario = []*unit.Case{
		{
			Name:   "area_success",
			Func:   unit.Func(area(3, 4)),
			Expect: unit.Returns(12, nil),
		},
		{
			Name:   "area_negative_height_success",
			Func:   unit.Func(area(-1, 4)),
			Expect: unit.Returns(0, errNegativeHeight),
		},
		{
			Name:   "area_negative_width_success",
			Func:   unit.Func(area(4, -1)),
			Expect: unit.Returns(0, errNegativeWidth),
		},
		{
			Name:   "area_fail",
			Func:   unit.Func(area(0, 0)),
			Expect: unit.Returns(0, nil),
		},
	}
)

var (
	errNegativeHeight error = errors.New("'height' can not be negative")
	errNegativeWidth  error = errors.New("'width' can not be negative")
)

func area(height, width int) (int, error) {
	if height < 0 {
		return 0, errNegativeHeight
	}
	if width < 0 {
		return 0, errNegativeWidth
	}
	result := height * width
	return result, nil
}

func TestMain(t *testing.T) {
	unit.Test(t, scenario)
}
```

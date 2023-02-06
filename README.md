## Welcome to Taste
*A taste of simplicity.*

taste is a simple (but powerful) table driven testing tool.

## Capabilities:

- HTTP Server Testing: 

It allows you to define a scenario of HTTP requests and expected responses, and will run the scenario and perform assertions to ensure that the server is behaving as expected. Check [example testing](example/server/test_case.go)


- Unit testing:

 It provides a set of functions and types for defining and running tests, and for asserting the expected results of those tests. The package is designed to be simple and flexible, and to make it easy to write and maintain unit tests. It supports table-driven testing, which allows you to define multiple tests in a single function using a slice of structs. It also supports custom expectations, which allow you to define the expected result of a test in a flexible way. 

### Example (HTTP Test):





---


## JSON body Validation

### Scheme

- The validation is done against a JSON validation scheme, which specifies the expected structure and content of the data.

- The validation scheme uses a specific syntax to specify the structure and content of a JSON object. The syntax is `"field_name|type": "value"`, where field_name is the name of the field, type is the data type of the field, and value is the expected value of the field.

- The validation scheme supports wildcards, which can be used to match any value or data type. The wildcard is denoted by *.

- The validate function returns an error if the data does not match the scheme, and returns nil if the data matches the scheme.

### Double bracket notation

```go
var (
	Scenario = []*taste.HTTPTestCase{
		{
			StoreResponse: true,
			Name:          "test_1",
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
			Name:          "test_2",
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
```

Double angle bracket notation (e.g. <<test_1.random>>) is used to reference the value of a field in a previous HTTP response.

For example, in the second test case (test_2), the request body includes the field `"id":"<<test_1.random>>"`. This means that the value of the id field should be the value of the random field in the response of the first test case (test_1).

This allows you to reuse values from previous responses in subsequent requests, which can be useful for testing scenarios where the server relies on data from previous requests.

Check out [this test file](example/double_angle_notation/main_test.go) for more complete example.

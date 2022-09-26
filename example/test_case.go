package example

import (
	"net/http"
	"taste"
)

var (
	TestUser *User = &User{ID: "718c9a02", Name: "john", Age: 20}
)

var (
	TestCases = []*taste.Case{
		{
			Name: "version_success",
			Request: &taste.Request{
				Method: http.MethodGet,
				URL:    "/api/v1/version",
			},
			Expect: &taste.Expect{
				Status:     http.StatusOK,
				BodyString: "v1.0.0",
			},
		},
		{
			Name: "users_success",
			Request: &taste.Request{
				Method: http.MethodGet,
				URL:    "/api/v1/users",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(Users),
			},
		},
		{
			Name: "user_fail",
			Request: &taste.Request{
				Method: http.MethodGet,
				URL:    "/api/v1/user?name=corey",
			},
			Expect: &taste.Expect{
				Status: http.StatusNotFound,
				Body:   []byte("404 page not found"),
			},
		},
		{
			Name: "user_get_eco_success",
			Request: &taste.Request{
				Method: http.MethodGet,
				URL:    "/api/v1/user?name=eco",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(Users[0]),
			},
		},
		{
			Name: "user_get_any_success",
			Request: &taste.Request{
				Method: http.MethodGet,
				URL:    "/api/v1/user?name=any",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(Users[1]),
				Header: http.Header{
					"Content-Type": []string{"application/json; charset=utf-8"},
				},
			},
		},
		{
			Name: "user_new_success",
			Request: &taste.Request{
				Method: http.MethodPost,
				URL:    "/api/v1/user/new",
				Body:   MarshalDiscardError(TestUser),
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(append(Users, TestUser)),
				Header: http.Header{
					"Content-Type": []string{"application/json; charset=utf-8"},
				},
			},
		},
	}
)

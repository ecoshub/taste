package example

import (
	"net/http"

	"github.com/ecoshub/taste/server"
)

var (
	TestUser *User = &User{ID: "718c9a02", Name: "john", Age: 20}
)

var (
	Scenario = []*server.Case{
		{
			Name: "version_success",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/version",
			},
			Expect: &server.Expect{
				Status:     http.StatusOK,
				BodyString: "v1.0.0",
			},
		},
		{
			Name: "users_success",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/users",
			},
			Expect: &server.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(Users),
			},
		},
		{
			Name: "user_fail",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=corey",
			},
			Expect: &server.Expect{
				Status: http.StatusNotFound,
				Body:   []byte("404 page not found"),
			},
		},
		{
			Name: "user_get_eco_success",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=eco",
			},
			Expect: &server.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(Users[0]),
			},
		},
		{
			Name: "user_get_any_success",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=any",
			},
			Expect: &server.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(Users[1]),
				Header: http.Header{
					"Content-Type": []string{"application/json; charset=utf-8"},
				},
			},
		},
		{
			Name: "user_new_success",
			Request: &server.Request{
				Method: http.MethodPost,
				Path:   "/api/v1/user/new",
				Body:   MarshalDiscardError(TestUser),
			},
			Expect: &server.Expect{
				Status: http.StatusOK,
				Body:   MarshalDiscardError(append(Users, TestUser)),
				Header: http.Header{
					"Content-Type": []string{"application/json; charset=utf-8"},
				},
			},
		},
	}
)

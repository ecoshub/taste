package example

import (
	"net/http"

	"github.com/ecoshub/taste/server"
)

var (
	TestUser []byte = []byte(`{"id":"718c9a02","name":"john","age":20}`)
)

var (
	Scenario = []*server.Case{
		{
			Name:        "version_success",
			CheckHeader: true,
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/version",
			},
			Expect: &server.Expect{
				Status:     http.StatusOK,
				Header:     server.HeaderContentPlainText,
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
				BodyString: `
					[
						{"id":"a4fb4201","name":"eco","age":30},
						{"id":"43bd1a0d","name":"any","age":29}
					]`,
			},
		},
		{
			Name: "user_fail",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=corey",
			},
			Expect: &server.Expect{
				Status:     http.StatusNotFound,
				BodyString: "404 page not found",
			},
		},
		{
			Name: "user_get_eco_success",
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=eco",
			},
			Expect: &server.Expect{
				Status:     http.StatusOK,
				BodyString: `{"id":"a4fb4201","name":"eco","age":30}`,
			},
		},
		{
			Name:        "user_get_any_success",
			CheckHeader: true,
			Request: &server.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=any",
			},
			Expect: &server.Expect{
				Status:     http.StatusOK,
				BodyString: `{"id":"43bd1a0d","name":"any","age":29}`,
				Header:     server.HeaderContentApplicationJSON,
			},
		},
		{
			Name:        "user_new_success",
			CheckHeader: true,
			Request: &server.Request{
				Method:     http.MethodPost,
				Path:       "/api/v1/user/new",
				BodyString: `{"id":"718c9a02","name":"john","age":20}`,
			},
			Expect: &server.Expect{
				Status: http.StatusOK,
				BodyString: `
					[
						{"id":"a4fb4201","name":"eco","age":30},
						{"id":"43bd1a0d","name":"any","age":29},
						{"id":"718c9a02","name":"john","age":20}
					]`,
				Header: server.HeaderContentApplicationJSON,
			},
		},
	}
)

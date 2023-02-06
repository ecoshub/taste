package example

import (
	"net/http"

	"github.com/ecoshub/taste"
)

var (
	Scenario = []*taste.HTTPTestCase{
		{
			Name: "version_success",
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/version",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Header: taste.HeaderContentPlainText,
				Body:   "v1.0.0",
			},
		},
		{
			Name: "users_success",
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/users",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body: `
					[
						{"id":"a4fb4201","name":"eco","age":30},
						{"id":"43bd1a0d","name":"any","age":29}
					]`,
			},
		},
		{
			Name: "user_fail",
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=corey",
			},
			Expect: &taste.Expect{
				Status: http.StatusNotFound,
				Body:   "404 page not found",
			},
		},
		{
			Name: "user_get_eco_success",
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=eco",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   `{"id":"a4fb4201","name":"eco","age":30}`,
			},
		},
		{
			Name:        "user_get_any_success",
			CheckHeader: true,
			Request: &taste.Request{
				Method: http.MethodGet,
				Path:   "/api/v1/user?name=any",
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body:   `{"id":"43bd1a0d","name":"any","age":29}`,
				Header: taste.HeaderContentApplicationJSON,
			},
		},
		{
			Name:        "user_new_success",
			CheckHeader: true,
			Request: &taste.Request{
				Method: http.MethodPost,
				Path:   "/api/v1/user/new",
				Body:   `{"id":"718c9a02","name":"john","age":20}`,
			},
			Expect: &taste.Expect{
				Status: http.StatusOK,
				Body: `
					[
						{"id":"a4fb4201","name":"eco","age":30},
						{"id":"43bd1a0d","name":"any","age":29},
						{"id":"718c9a02","name":"john","age":20}
					]`,
				Header: taste.HeaderContentApplicationJSON,
			},
		},
	}
)

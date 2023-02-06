package taste

import (
	"errors"
	"testing"
)

var (
	scenario = []*UnitTestCase{
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(``),
				[]byte(``),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{}`),
				[]byte(`{}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{}`),
				[]byte(``),
			)),
			Expect: Returns(errors.New("expected something but got nothing (got nil body)")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name":"eco","age":30,"items":["cellphone","table"]}`),
				[]byte(`{"name":"eco","age":30,"items":["cellphone","table","chair"]}`),
			)),
			Expect: Returns(errors.New("unexpected path. path: [items 2]")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name":"eco"}`),
				[]byte(`{"name":"test"}`),
			)),
			Expect: Returns(errors.New("value expectation failed. expected value: 'eco', got value: 'test', path: '[name]'")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name":"eco"}`),
				[]byte(`{"name":"eco"}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name":"eco","age":30}`),
				[]byte(`{"name":"eco"}`),
			)),
			Expect: Returns(errors.New("field is required. field: 'age'")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name":"eco","age":30}`),
				[]byte(`{"name":"eco","age":"30"}`),
			)),
			Expect: Returns(errors.New("type expectation failed. expected type: 'int', got type: 'string', path: '[age]'")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"eco","age|int":30}`),
				[]byte(`{"name":"eco","age":"30"}`),
			)),
			Expect: Returns(errors.New("type expectation failed. expected type: 'int', got type: 'string', path: '[age]'")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"eco","age|int":30}`),
				[]byte(`{"name":"eco","age":30.0}`),
			)),
			Expect: Returns(errors.New("type expectation failed. expected type: 'int', got type: 'float', path: '[age]'")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"eco","age|int":30}`),
				[]byte(`{"name":"eco","age":30}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"eco","age|int":30}`),
				[]byte(`{"name": "eco", "age" :30}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"eco","age|int":30}`),
				[]byte(`{"name": "eco", "age" :30, "extra":true}`),
			)),
			Expect: Returns(errors.New("unexpected path. path: [extra]")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"eco","age|int":30}`),
				[]byte(`{"age" :30, "name": "eco"}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"*","age|int":"*"}`),
				[]byte(`{"age" :30, "name": "eco"}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"*","age|int":"*"}`),
				[]byte(`{"age" :"emre", "name": 30}`),
			)),
			Expect: Returns(errors.New("type expectation failed. expected type: 'string', got type: 'int', path: '[name]'")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"*","age|int":"*"}`),
				[]byte(`{"age":72,"name":"test"}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"*","age|int":"30", "*employed|boolean":"*"}`),
				[]byte(`{"age":30,"name":"eco"}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"*","age|int":"30", "*employed|boolean":"*"}`),
				[]byte(`{"age":30,"name":"eco","employed":"yes"}`),
			)),
			Expect: Returns(errors.New("type expectation failed. expected type: 'boolean', got type: 'string', path: '[employed]'")),
		},
		{
			Name: "validation_fuzzy_object",
			Func: Func(validate(
				[]byte(`{"name|string":"*","age|int":"30", "*employed|boolean":"*"}`),
				[]byte(`{"age":30,"name":"eco","employed":false}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`[]`),
				[]byte(`[]`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`[]`),
				[]byte(``),
			)),
			Expect: Returns(errors.New("expected something but got nothing (got nil body)")),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(``),
				[]byte(`[]`),
			)),
			Expect: Returns(errors.New("no expectation specified but got response body. body: '[]'")),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":30}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":"30"}
				]}`),
			)),
			Expect: Returns(errors.New("type expectation failed. expected type: 'int', got type: 'string', path: '[items 0 age]'")),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":30}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":30}
				]}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":"*"}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":31},
					{"name":"eco2","age":32},
					{"name":"eco3","age":33}
				]}`),
			)),
			Expect: Returns(errors.New("unexpected path. path: [items 1 name]")),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":"*","games|array":[
						"headball2",
						"*"
					]}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":31,"games":"test"}
				]}`),
			)),
			Expect: Returns(errors.New("type expectation failed. expected type: 'array', got type: 'string', path: '[items 0 games]'")),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":"*","games|array":[
						"headball2",
						"*"
					]}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":31,"games":[]}
				]}`),
			)),
			Expect: Returns(errors.New("error: array is empty error_code: 02. path: [items 0 games 0]")),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":"*","games|array":[
						"headball2",
						"*"
					]}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":31,"games":[
						"headball2"
					]}
				]}`),
			)),
			Expect: Returns(errors.New("error: index out of range error_code: 07. path: [items 0 games 1]")),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":"*","games|array":[
						"headball2",
						"*"
					]}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":31,"games":[
						"headball2",
						"eco"
					]}
				]}`),
			)),
			Expect: Returns(nil),
		},
		{
			Name: "validation_fuzzy_array",
			Func: Func(validate(
				[]byte(`{"items|array":[
					{"name|string":"*","age|int":"*","games|array":[
						"headball2",
						"*"
					]}
				]}`),
				[]byte(`{"items":[
					{"name":"eco1","age":31,"games":[
						"headball2",
						"eco",
						"eco"
					]}
				]}`),
			)),
			Expect: Returns(errors.New("unexpected path. path: [items 0 games 2]")),
		},
		{
			Name: "validation_null_values",
			Func: Func(validate(
				[]byte(`{
					"name":"",
					"age":0,
					"body":{
						"test":"",
						"emre":null,
						"run":false,
						}
					}`),
				[]byte(`{"name":"", "age":0, "body":{"test":"","emre":null,"run":false}}`),
			)),
			Expect: Returns(nil),
		},
	}
)

func TestValidation(t *testing.T) {
	Run(t, scenario)
}

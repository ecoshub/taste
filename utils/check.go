package utils

import (
	"reflect"
	"testing"
)

var (
	TestLogEnable bool
)

func CheckEqualWithNilError(t *testing.T, field string, err error, got, expect interface{}) {
	CheckExpectError(t, field+"-error", err, nil)
	CheckEqual(t, field, got, expect)
}

func CheckNotEqual(t *testing.T, field string, got, expect interface{}) {
	if check(t, field, got, expect) {
		FailNot(t, field, got, expect)
	}
}

func CheckEqual(t *testing.T, field string, got, expect interface{}) {
	if !check(t, field, got, expect) {
		Fail(t, field, got, expect)
	}
}

func check(t *testing.T, field string, got, expect interface{}) bool {
	return reflect.DeepEqual(got, expect)
}

func CheckExpectError(t *testing.T, field string, got, expect error) {
	gotErrMsg := "<nil>"
	if got != nil {
		gotErrMsg = got.Error()
	}

	expectErrMsg := "<nil>"
	if expect != nil {
		expectErrMsg = expect.Error()
	}

	if gotErrMsg != expectErrMsg {
		Fail(t, field, got, expect)
		return
	}
}

func FailNot(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf("'%v' unexpected result. got: '%v'%T, not expected: '%v'(%T)", field, got, got, expect, expect)
}

func Fail(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf("'%v' unexpected result. got: '%v'(%T), expected: '%v'(%T)", field, got, got, expect, expect)
}

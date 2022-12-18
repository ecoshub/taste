package utils

import (
	"reflect"
	"testing"
)

const (
	expectScheme    = "'%v' unexpected result. got: '%v'(%T), expected: '%v'(%T)"
	expectNotScheme = "'%v' unexpected result. got: '%v'(%T), not expected: '%v'(%T)"
)

var (
	TestLogEnable bool
)

func CheckEqualWithNilError(t *testing.T, field string, err error, got, expect interface{}) {
	CheckExpectError(t, field+"-error", err, nil)
	CheckEqual(t, field, got, expect)
}

func CheckNotEqual(t *testing.T, field string, got, expect interface{}) {
	if Check(got, expect) {
		FailNot(t, field, got, expect)
	}
}

func CheckEqual(t *testing.T, field string, got, expect interface{}) {
	if !Check(got, expect) {
		Fail(t, field, got, expect)
	}
}

func Check(got, expect interface{}) bool {
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

func CheckEqualLog(t *testing.T, field string, got, expect interface{}) bool {
	if Check(got, expect) {
		return false
	}
	t.Logf(expectScheme, field, got, got, expect, expect)
	return true
}

func FailNot(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf(expectNotScheme, field, got, got, expect, expect)
}

func Fail(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf(expectScheme, field, got, got, expect, expect)
}

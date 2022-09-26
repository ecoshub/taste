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
	LogChecking(t, "Equal", field, got, expect)
	if reflect.DeepEqual(got, expect) {
		FailNot(t, field, got, expect)
	}
	LogDone(t, "Equal", field)
}

func CheckEqual(t *testing.T, field string, got, expect interface{}) {
	LogChecking(t, "Equal", field, got, expect)
	if !reflect.DeepEqual(got, expect) {
		Fail(t, field, got, expect)
	}
	LogDone(t, "Equal", field)
}

func CheckExpectError(t *testing.T, field string, got, expect error) {
	LogChecking(t, "Error expect", field, got, expect)
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

	LogDone(t, "Error expect", field)
}

func FailNot(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf("'%v' unexpected result. got: '%v'%T, not expected: '%v'(%T)", field, got, got, expect, expect)
}

func Fail(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf("'%v' unexpected result. got: '%v'(%T), expected: '%v'(%T)", field, got, got, expect, expect)
}

func LogChecking(t *testing.T, functionName, field string, got, expect interface{}) {
	if !TestLogEnable {
		return
	}
	t.Logf("\t> checking func: %v field: %v. got: %v, expected: %v\n", functionName, field, got, expect)
}

func LogDone(t *testing.T, functionName, field string) {
	if !TestLogEnable {
		return
	}
	t.Logf("\t> done func: %v field: %v.\n", functionName, field)
}

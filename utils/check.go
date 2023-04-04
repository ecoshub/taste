// Package utils provides utility functions for testing.

package utils

import (
	"reflect"
	"testing"
)

// TestLogEnable specifies whether to enable logging of test results.
var TestLogEnable bool

// CheckEqualWithNilError checks that two values are equal and the error is nil.
func CheckEqualWithNilError(t *testing.T, field string, err error, got, expect interface{}) {
	CheckExpectError(t, field+"-error", err, nil)
	CheckEqual(t, field, got, expect)
}

// CheckNotEqual checks that two values are not equal.
func CheckNotEqual(t *testing.T, field string, got, expect interface{}) {
	if check(t, field, got, expect) {
		failNot(t, field, got, expect)
	}
}

// CheckEqual checks that two values are equal.
func CheckEqual(t *testing.T, field string, got, expect interface{}) {
	if !check(t, field, got, expect) {
		fail(t, field, got, expect)
	}
}

// check checks that two values are equal using reflection.
func check(t *testing.T, field string, got, expect interface{}) bool {
	return reflect.DeepEqual(got, expect)
}

// CheckExpectError checks that two errors are equal.
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
		fail(t, field, got, expect)
		return
	}
}

// failNot logs a failure message for CheckNotEqual.
func failNot(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf("'%s' unexpected result. got: '%v' (%T), not expected: '%v' (%T)", field, got, got, expect, expect)
}

// fail logs a failure message for CheckEqual.
func fail(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf("'%s' unexpected result. got: '%v' (%T), expected: '%v' (%T)", field, got, got, expect, expect)
}

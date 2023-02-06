package taste

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

func check(got, expect interface{}) bool {
	return reflect.DeepEqual(got, expect)
}

func checkExpectError(t *testing.T, field string, got, expect error) {
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

func checkEqualLog(t *testing.T, field string, got, expect interface{}) bool {
	if check(got, expect) {
		return false
	}
	t.Logf(expectScheme, field, got, got, expect, expect)
	return true
}

func fail(t *testing.T, field string, got, expect interface{}) {
	t.Fatalf(expectScheme, field, got, got, expect, expect)
}

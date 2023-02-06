package taste

import (
	"fmt"
	"reflect"
	"testing"
)

type UnitTestCase struct {
	// Name name of the test case
	Name string

	// Func function to test
	Func []interface{}

	// Expect expected response
	Expect []interface{}

	// OnlyRunThis if this flag enabled, test system will run only this
	// if multiple test case enabled this flag, system will run the first occurred test case only
	OnlyRunThis bool
}

// Func function wrapper function
func Func(i ...interface{}) []interface{} { return i }

// Returns return values wrapper function
func Returns(i ...interface{}) []interface{} { return i }

// Run run the test cases
func Run(t *testing.T, scenario []*UnitTestCase) {
	c, ok := hasOnlyRunMe(scenario)
	if ok {
		c.Test(t)
		return
	}
	for _, c := range scenario {
		c.Test(t)
	}
}

// hasOnlyRunMe find the case that has OnlyRunThis flag enabled
func hasOnlyRunMe(scenario []*UnitTestCase) (*UnitTestCase, bool) {
	for _, c := range scenario {
		if c.OnlyRunThis {
			return c, true
		}
	}
	return nil, false
}

func (c *UnitTestCase) Test(t *testing.T) {
	t.Run(c.Name, func(t *testing.T) {
		if c.Func == nil {
			t.Fatalf("error at test case '%s'. Field required 'Func'", c.Name)
		}

		if c.Expect == nil {
			t.Fatalf("error at test case '%s'. Field required 'Expect'", c.Name)
		}

		if len(c.Expect) != len(c.Func) {
			t.Fatalf("test function returns %d value but expected %d value.", len(c.Func), len(c.Expect))
		}

		// compare all responses
		for i := range c.Func {
			expected := c.Func[i]
			got := c.Expect[i]

			if isNil(expected) && isNil(got) {
				continue
			}

			if !reflect.DeepEqual(expected, got) {
				t.Fatalf("expectations does not satisfied. expected: '%v' (%T), got: '%v' (%T)", expected, expected, got, got)
			}
		}
	})
}

func isNil(i interface{}) bool {
	return fmt.Sprint(i) == "<nil>"
}

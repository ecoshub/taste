package unit

import (
	"fmt"
	"reflect"
	"testing"
)

// Case represents a single test case
type Case struct {
	Name        string        // Name of the test case
	OnlyRunThis bool          // If true, only run this test case and skip others
	Func        []interface{} // Arguments to the function being tested
	Expect      []interface{} // Expected return values from the function being tested
}

// Func is a helper function to pass arguments to a test case
func Func(params ...interface{}) []interface{} { return params }

// Returns is a helper function to pass expected return values to a test case
func Returns(params ...interface{}) []interface{} { return params }

// Test runs the provided test cases
func Test(t *testing.T, scenario []*Case) {
	for _, c := range scenario {
		if c.OnlyRunThis {
			c.Test(t)
			return
		}
	}
	// Otherwise, run all test cases in the provided order
	for _, c := range scenario {
		c.Test(t)
	}
}

func (c *Case) Test(t *testing.T) {
	t.Run(c.Name, func(t *testing.T) {
		if c.Func == nil {
			t.Fatalf("nil function")
		}

		if c.Expect == nil {
			t.Fatalf("nil expectations. got: %v", c.Func)
		}

		if len(c.Expect) != len(c.Func) {
			t.Fatalf("test function returns %d value but expected %d value.", len(c.Func), len(c.Expect))
		}

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

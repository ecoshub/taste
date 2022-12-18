package unit

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	Name        string
	Func        []interface{}
	Expect      []interface{}
	OnlyRunThis bool
}

func Func(i ...interface{}) []interface{}    { return i }
func Returns(i ...interface{}) []interface{} { return i }

func Test(t *testing.T, scenario []*Case) {
	c, ok := hasOnlyRunMe(scenario)
	if ok {
		c.Test(t)
		return
	}
	for _, c := range scenario {
		c.Test(t)
	}
}

func hasOnlyRunMe(scenario []*Case) (*Case, bool) {
	for _, c := range scenario {
		if c.OnlyRunThis {
			return c, true
		}
	}
	return nil, false
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

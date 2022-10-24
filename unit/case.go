package unit

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

type Case struct {
	FieldName       string
	FieldGot        []interface{}
	FieldExpect     []interface{}
	FieldExpectFail bool
}

func NewCase() *Case {
	return &Case{}
}

func (c *Case) Name(name string) *Case {
	c.FieldName = name
	return c
}

func (c *Case) Run(i ...interface{}) *Case {
	c.FieldGot = i
	return c
}

func (c *Case) Expect(i ...interface{}) *Case {
	c.FieldExpect = i
	return c
}

func (c *Case) ExpectFail() *Case {
	c.FieldExpectFail = true
	return c
}

func main() {

	// for _, c := range cases {
	// 	fmt.Println(c.Name)
	// 	fmt.Println(c.run)

	// }
}

func sum(a, b int) int {
	return a + b
}

func getMarried(name string, age int, married bool) (*Person, error) {
	if name == "eco" && !married {
		return nil, errors.New("of course you are married")
	}
	if married {
		return nil, errors.New("you'r already married")
	}
	if name != "eco" {
		if age >= 30 {
			return &Person{Name: "person to marry", Age: 25}, nil
		}
	}
	return nil, errors.New("you have more time")
}

func (c *Case) Test(t *testing.T) {
	t.Run(c.FieldName, func(t *testing.T) {
		if c.FieldGot == nil {
			t.Fatalf("nil function")
		}

		if c.FieldExpect == nil {
			t.Fatalf("nil expectations. got: %v", c.FieldGot)
		}

		if len(c.FieldExpect) != len(c.FieldGot) {
			t.Fatalf("function returns and expected value counts are not same")
		}

		for i := range c.FieldGot {
			expected := c.FieldGot[i]
			got := c.FieldExpect[i]

			if isNil(expected) && isNil(got) {
				continue
			}

			if !reflect.DeepEqual(expected, got) {
				if !c.FieldExpectFail {
					t.Fatalf("expectations does not satisfied. expected: '%v' (%T), got: '%v' (%T)", expected, expected, got, got)
				}
			} else {
				if c.FieldExpectFail {
					t.Fatalf("expected to failed but pass. expected: '%v' (%T), got: '%v' (%T)", expected, expected, got, got)
				}
			}
		}
	})
}

func isNil(i interface{}) bool {
	return fmt.Sprint(i) == "<nil>"
}

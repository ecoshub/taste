package unit

// import (
// 	"errors"
// 	"fmt"
// 	"reflect"
// 	"testing"
// )

// var (
// 	cases = []*Case{
// 		{
// 			FieldName:   "sum_2_3_success",
// 			FieldGot:    Wrap_(sum(2, 3)),
// 			FieldExpect: Wrap_(5),
// 		},
// 		{
// 			FieldName:       "sum_5_8_fail",
// 			FieldGot:        Wrap_(sum(2, 3)),
// 			FieldExpect:     Wrap_(5),
// 			FieldExpectFail: true,
// 		},
// 		{
// 			FieldName:   "marry_fail",
// 			FieldGot:    Wrap_(getMarried("alican", 28, false)),
// 			FieldExpect: Wrap_(nil, "you have more time"),
// 		},
// 	}
// )

// func Wrap_(i ...interface{}) []interface{} { return i }

// func TestUnit(t *testing.T) {
// 	NewCase().
// 		Name("sum_5_5_success").
// 		Expect(10).
// 		Run(sum(5, 5)).
// 		Test(t)

// 	NewCase().
// 		Name("sum_5_5_success").
// 		Expect(9).
// 		Run(sum(5, 5)).
// 		ExpectFail().
// 		Test(t)

// 	NewCase().
// 		Name("marry_success").
// 		Expect(nil, errors.New("you have more time")).
// 		Run(getMarried("alican", 28, false)).
// 		Test(t)

// 	NewCase().
// 		Name("marry_fail").
// 		Expect(nil, errors.New("of course you are married")).
// 		Run(getMarried("eco", 28, false)).
// 		Test(t)

// 	NewCase().
// 		Name("marry_fail").
// 		Expect(&Person{Name: "person to marry", Age: 25}, nil).
// 		Run(getMarried("alican", 35, false)).
// 		Test(t)
// }

package main

import (
	"testing"

	"github.com/ecoshub/taste"
)

var (
	scenario = []*taste.UnitTestCase{
		{
			Name:   "area_success",
			Func:   taste.Func(area(3, 4)),
			Expect: taste.Returns(12, nil),
		},
		{
			Name:   "area_negative_height_success",
			Func:   taste.Func(area(-1, 4)),
			Expect: taste.Returns(0, errNegativeHeight),
		},
		{
			Name:   "area_negative_width_success",
			Func:   taste.Func(area(4, -1)),
			Expect: taste.Returns(0, errNegativeWidth),
		},
		{
			Name:   "area_negative_height_and_width_success",
			Func:   taste.Func(area(-1, -1)),
			Expect: taste.Returns(0, errNegativeHeight),
		},
		{
			Name:   "area_fail",
			Func:   taste.Func(area(0, 0)),
			Expect: taste.Returns(0, nil),
		},
	}
)

func TestMain(t *testing.T) {
	taste.Run(t, scenario)
}

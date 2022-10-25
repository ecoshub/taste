package main

import (
	"testing"

	"github.com/ecoshub/taste/unit"
)

var (
	scenario = []*unit.Case{
		{
			Name:   "area_success",
			Func:   unit.Func(area(3, 4)),
			Expect: unit.Returns(12, nil),
		},
		{
			Name:   "area_negative_height_success",
			Func:   unit.Func(area(-1, 4)),
			Expect: unit.Returns(0, errNegativeHeight),
		},
		{
			Name:   "area_negative_width_success",
			Func:   unit.Func(area(4, -1)),
			Expect: unit.Returns(0, errNegativeWidth),
		},
		{
			Name:   "area_negative_height_and_width_success",
			Func:   unit.Func(area(-1, -1)),
			Expect: unit.Returns(0, errNegativeHeight),
		},
		{
			Name:   "area_fail",
			Func:   unit.Func(area(0, 0)),
			Expect: unit.Returns(0, nil),
		},
	}
)

func TestMain(t *testing.T) {
	unit.Test(t, scenario)
}

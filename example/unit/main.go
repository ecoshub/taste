package main

import "errors"

var (
	errNegativeHeight error = errors.New("'height' can not be negative")
	errNegativeWidth  error = errors.New("'width' can not be negative")
)

func area(height, width int) (int, error) {
	if height < 0 {
		return 0, errNegativeHeight
	}
	if width < 0 {
		return 0, errNegativeWidth
	}
	result := height * width
	return result, nil
}

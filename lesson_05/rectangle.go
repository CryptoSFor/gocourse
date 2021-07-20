package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() (float64, error) {
	if r.width > 0 && r.height > 0 {
		return r.height * r.width, nil
	} else {
		return 0, errors.New("incorrect input")
	}
}

func (r Rectangle) Perimeter() (float64, error) {
	if r.width > 0 && r.height > 0 {
		return 2 * (r.width + r.height), nil
	} else {
		return 0, errors.New("incorrect input")
	}
}

func (r Rectangle) String() string {
	if r.width > 0 && r.height > 0 {
		return fmt.Sprintf("Rectangle with height %.2f and width %.2f", r.height, r.width)
	} else {
		return fmt.Sprintf("Incorrect input; height = %.2f and width = %.2f", r.height, r.width)
	}
}

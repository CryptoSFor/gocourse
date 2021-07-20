package main

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() (float64, error) {
	if c.radius > 0 {
		return math.Pow(c.radius, 2) * math.Pi, nil
	} else {
		return 0, errors.New("Incorrect input")
	}
}

func (c Circle) Perimeter() (float64, error) {
	if c.radius > 0 {
		return 2 * c.radius * math.Pi, nil
	} else {
		return 0, errors.New("Incorrect input")
	}
}

func (c Circle) String() string {
	if c.radius > 0 {
		return fmt.Sprintf("Circle: radius %.2f", c.radius)
	} else {
		return fmt.Sprintf("Incorrect input; radius = %.2f", c.radius)
	}
}

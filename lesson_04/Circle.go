package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", c.radius)
}

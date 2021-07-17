package main

import "fmt"

func main() {
	c := Circle{radius: 8}
	r := Rectangle{
		height: 9,
		width:  3,
	}
	DescribeShape(c)
	DescribeShape(r)
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

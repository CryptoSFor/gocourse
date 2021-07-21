package main

import "fmt"

func main() {
	c := Circle{radius: 9}
	r := Rectangle{
		height: 9,
		width:  3,
	}
	DescribeShape(c)
	DescribeShape(r)
}

func DescribeShape(s Shape) {
	fmt.Println(s)

	area, err_area := s.Area()
	if err_area == nil {
		fmt.Printf("Area: %.2f\n", area)
	}

	per, err_perim := s.Perimeter()
	if err_perim == nil {
		fmt.Printf("Perimeter: %.2f\n", per)
	}
}

package main

import (
	"fmt"
	"main/fib_sequence"
)

func main() {
	fmt.Println("Fib nums:")
	fib_sequence.FibNums(6)
	defer fmt.Println("Complete")
}

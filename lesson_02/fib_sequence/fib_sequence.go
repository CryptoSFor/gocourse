package fib_sequence

import "fmt"

func FibNums(n int) {
	var a, b int = 1, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

package fibonacci

import "fmt"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func Print(n int) {
	f := Fibonacci()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", f())
	}
}

package main

import "fmt"

func fibonacci1 (n int) int {
	
	var a1, a2 = 0, 1

	if n == 1 {
		return a1
	}
	if n == 2 {
		return a2
	}
	
	var fib int
	
	for i := 2; i < n; i++ {
		fib = a1 + a2
		a1 = a2
		a2 = fib
	}
	return fib
}

func main() {
	fmt.Println(fibonacci1(13))
}
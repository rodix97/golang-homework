package main

import "fmt"

func fibonacci2 (n int) int {
	
	var a1, a2 = 0, 1
	
	if n == 1 {
		return a1
	}

	if n == 2 {
		return a2
	}

	p1 := fibonacci2(n-1)
	p2 := fibonacci2(n-2)
	
	return p1 + p2
}

func main() {
	fmt.Println(fibonacci2(9))
}
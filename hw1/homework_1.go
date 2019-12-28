package main

import "fmt"

func multiply(a int, b int) int {
	otvet := 0
	for i := 0; i < b; i++ {
		otvet += a
	}

	return otvet
}

func main() {
	fmt.Println(multiply(5,9))
}
package main

import "fmt"

func unique_count (b []int) int {
	
	var c []int
	
	for i := 0; i < len(b); i++ {
		k := 0
		for j := 0; j < len(c); j++ {
			if b[i] != c[j] {
				k++
			}
		}
		if k == len(c) {
			c = append(c,b[i])
		}
	}
	return len(c)
}

func main() {
	a := [...]int{1, 2, 3, 4, 1, 2, 2, 3, 2}
	fmt.Println(unique_count(a[:]))
}
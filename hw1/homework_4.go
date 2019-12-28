package main

import "fmt"

func bubble_sort (b []int) {
	
	k := 1
	n := len(b)

	for i := n; i >= 0; i-- {
		if k > 0 {
			k = 0
		} else {
			break
		}
		for j := 0; j < i-1; j++ {
			if b[j] > b[j+1] {
				bub := b[j]
				b[j] = b[j+1]
				b[j+1] = bub
				k++
			}
		}
	}

}

func main() {
	a := [10]int{25, 45, 6, 13, 3, 35, 17, 9, 24, 5}
	bubble_sort(a[:])
	fmt.Println(a)
}
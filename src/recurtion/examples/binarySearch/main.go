package main

import (
	"fmt"
)

func binarySearch(v []int, a, b, c int) int {
	if a >= b {
		if v[a] == c {
			return a
		}else {
			return -1
		}
	}

	if v[(a+b)/2] == c {
		return (a+b)/2
	}
	if v[(a+b)/2] < c {
		return binarySearch(v, (a+b)/2+1, b, c)
	}else {
		return binarySearch(v, a, (a+b)/2 -1, c)
	}	
}

func main() {
	v := make([]int, 0, 100)
	for i := range 100 {
		v = append(v, i)
	}
	fmt.Printf("Found %d\n", binarySearch(v, 5, 20, 8))
}
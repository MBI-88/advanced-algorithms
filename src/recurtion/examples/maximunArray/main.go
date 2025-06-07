package main

import (
	"fmt"
	"math"
)

func iterMaximun(v []int, a, b int) int {
	var (
		maximun = math.MinInt
	)

	for _, i := range v[a:b+1] {
		if i > maximun {
			maximun = i
		}
	}
	return maximun
}

func recursionMaximun(v []int, a, b int) int {
	var (
		left int 
		right int
	)
	if a == b {
		return v[a]
	}else if a < b {
		left = recursionMaximun(v, a, (a+b)/2)
		right  = recursionMaximun(v, (a+b)/2 + 1, b)
	}
	if left > right {
		return left
	}
	return right
}



func main() {
	array := []int{1,3,8,10,20,2,8,12,78,100}
	fmt.Printf("Maximun between position %d and %d is %d\n", 3, 6, recursionMaximun(array, 3, 6))
	fmt.Printf("Maximun between position %d and %d is %d\n", 3, 6, iterMaximun(array, 3, 6))
}
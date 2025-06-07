package main

import "fmt"

func mochila(b, p []int, T int) int {
	var (
		result int
		backtrack func(index, v, c int)
	)

	backtrack = func(index, v,  c int) {
		if v > result && c > 0 {
			result = v
		}else if c <= 0 {
			return
		}
		
		for i := index; i < len(b); i++ {
			backtrack(i+1, v + b[i], c - p[i])
		}
	}

	backtrack(0, 0, T)
	return result
}

func main() {
	b := []int{10,5, 8, 100}
	p := []int{5, 2, 1, 8}
	T := 10

	fmt.Printf("Benefice %d\n", mochila(b, p, T))

}
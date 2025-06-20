package main

import (
	"flag"
	"fmt"
)

var (
	p []int
	N *int
)

func init() {
	N = flag.Int("n", 5, "permutation number")
}

func permutation(n int) {
	var aux int
	
	if n == 0 {
		for i := 0; i < *N; i++ {
			fmt.Printf("%d ", p[i])
		}
		println()
	} else {
		for i := 0; i < n; i++ {
			aux = p[i]
			p[i] = p[n-1]
			p[n - 1] = aux
			permutation(n - 1)
			aux = p[i]
			p[i] = p[n-1]
			p[n-1] = aux
		}
	}
}

func main() {
	flag.Parse()

	p = make([]int, 0, *N+1)
	for i := 1; i <= *N; i++ {
		p = append(p, i)
	}
	
	permutation(*N)
}

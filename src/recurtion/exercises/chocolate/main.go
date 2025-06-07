package main

import (
	"fmt"
	"sort"
)

func chocolateCut(M, N int, cuts []int) int {
	if M < 2 || M > 1000 || N < 2 || N > 1000 {
		return 0
	}
	var (
		totalCost    int
		V            = cuts[:M-1]
		H            = cuts[N+1:]
		backtracking func(int, int)
		hSegment     = N+1
		vSegment     = M+1
	)
	sort.Sort(sort.Reverse(sort.IntSlice(V)))
	sort.Sort(sort.Reverse(sort.IntSlice(H)))

	backtracking = func(m, n int) {
		if m > 0 && n > 0 {
			if V[m] > H[n] {
				totalCost += V[m] * hSegment
				vSegment--
				backtracking(m-1, n)
			} else {
				totalCost += H[n] * vSegment
				hSegment--
				backtracking(m, n-1)
			}
		} else {
			for m > 0 {
				totalCost += V[m] * hSegment
				m--
			}
			for n > 0 {
				totalCost += H[n] * vSegment
				n--
			}
		}

	}

	backtracking(len(V)-1, len(H)-1)
	return totalCost
}

func main() {
	m := 6
	n := 4
	cuts := []int{2, 1, 3, 1, 4, 4, 1, 2}

	fmt.Printf("Coste total %d \n ", chocolateCut(m, n, cuts))

}

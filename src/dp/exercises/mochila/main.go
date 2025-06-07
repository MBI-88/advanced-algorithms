package main

import "fmt"

func knackpack(B, P []int, M int) int {
	var (
		m    = len(B) + 1
		K = make([][]int, m)
	)

	for i := range m {
		K[i] = make([]int, M+1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j <= M; j++ {
			if j >= P[i-1] &&  B[i-1]+K[i-1][j-P[i-1]] > K[i-1][j]  {
				K[i][j] = B[i-1]+K[i-1][j-P[i-1]]
			}else {
				K[i][j] = K[i-1][j]
			}

		}
	}
	
	return K[m-1][M]
}

func main() {
	b := []int{28,33,5,12,20}
	p := []int{2,5,3,6,1}
	T := 10

	fmt.Printf("Benefice gotton %d\n", knackpack(b, p, T))
}
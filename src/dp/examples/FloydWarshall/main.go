package main

import "fmt"

type tuple [2]int

// -1 = inf

func floydWarshall(N int, C map[tuple]int) [][]int {
	var P = make([][]int, N)

	for i := range N {
		P[i] = make([]int, N)
	}

	for k := range C {
		i, j := k[0], k[1]
		P[i][j] = C[k]
	}

	for k := range N {
		for i := range N {
			for j := range N {
				if P[i][k] != -1 && P[k][j] != -1 {
					if P[i][j] == -1 || P[i][j] > P[i][k]+P[k][j] {
						P[i][j] = P[i][k] + P[k][j]
					}
				}
			}
		}
	}
	return P
}

func main() {
	n := 5
	c := map[tuple]int{
		{0, 0}: 0,
		{0, 1}: 3,
		{0, 2}: -1,
		{0, 3}: 1,
		{0, 4}: 5,
		{1, 0}: -1,
		{1, 1}: 0,
		{1, 2}: 1,
		{1, 3}: -1,
		{1, 4}: -1,
		{2, 0}: -1,
		{2, 1}: -1,
		{2, 2}: 0,
		{2, 3}: 3,
		{2, 4}: 4,
		{3, 0}: -1,
		{3, 1}: 1,
		{3, 2}: 2,
		{3, 3}: 0,
		{3, 4}: -1,
		{4, 0}: -1,
		{4, 1}: -1,
		{4, 2}: -1,
		{4, 3}: 6,
		{4, 4}: 0,
	}

	fmt.Println(floydWarshall(n, c))
}

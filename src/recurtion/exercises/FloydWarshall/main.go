package main

import (
	"flag"
	"fmt"
	"math"
)

type tuple [2]int

func floydWarshall(start, end int, C map[tuple]int) int {
	var (
		searchNode func(int, int, int)
		memory     = make(map[int]bool, len(C)*2)
		bestRoute  int
		minimun    = math.MaxInt
	)

	searchNode = func(s, e, c int) {
		if _, ok := memory[s]; ok {
			return
		}

		for k := range C {
			if (k[0] == s) && C[k] > 0 {
				nextCost := c + C[k]
				if nextCost >= minimun {
					continue
				}
				memory[k[0]] = true
				defer delete(memory, k[0])

				if k[1] == e {
					minimun = nextCost

				} else {
					searchNode(k[1], e, nextCost)
				}

			}
		}

	}
	searchNode(start, end, bestRoute)
	return minimun
}

var (
	start *int
	end   *int
)

func init() {
	start = flag.Int("start", 0, "node start")
	end = flag.Int("end", 1, "node end")

}

func main() {
	flag.Parse()

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

	fmt.Println(floydWarshall(*start, *end, c))
}

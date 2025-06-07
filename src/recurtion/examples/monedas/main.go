package main

import (
	"fmt"
	"math"
)


func minExchange(v []int, C int) int {
	var (
		memo = make(map[int]int, C)
		track func([]int, int) int
	)

	track = func (vals []int, amount int ) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return math.MaxInt
		}

		if t,ok := memo[amount]; ok {
			return t
		}
		
		minChange := math.MaxInt
		for _, count := range vals {
			resp := track(vals, amount - count)
			if resp != math.MaxInt {
				minChange = int(math.Min(float64(minChange), float64(resp + 1)))
			}
		}

		memo[amount] = minChange
		return minChange
	}

	return track(v, C)
	
}

func main() {
	values := []int{5,3,2,7,1}
	exchange := 5

	fmt.Printf("Total monedas %d\n", minExchange(values, exchange))

}

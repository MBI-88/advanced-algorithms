package main

import (
	"fmt"
	"math"
)

func minExchange(vs []int, C int) int {
	n := len(vs)
	dp := make([][]*int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]*int, C+1)
	}

	zero := 0
	dp[0][0] = &zero

	for j := 1; j <= C; j++ {
		if j >= vs[0] && dp[0][j-vs[0]] != nil {
			val := *dp[0][j-vs[0]] + 1
			dp[0][j] = &val
		}
	}

	for i := 1; i < n; i++ {
		dp[i][0] = &zero
		for j := 1; j <= C; j++ {

			if j < vs[i] || dp[i][j-vs[i]] == nil {
				dp[i][j] = dp[i-1][j]
			}else if dp[i-1][j] != nil {
				temp1 := dp[i-1][j]
				temp2 := *dp[i][j-vs[i]] + 1
				result := int(math.Min(float64(*temp1), float64(temp2)))
				dp[i][j] = &result
			}else {
				 var v *int
				if dp[i][j-vs[i]] != nil {
					result := 1 + *dp[i][j-vs[i]]
					v = &result
				}
				dp[i][j] =  v
			}
		}
	}

	if dp[n-1][C] == nil {
		return -1
	}
	return *dp[n-1][C]
}

func main() {
	values := []int{2,10,1}
	exchange := 8

	fmt.Printf("Total monedas %d\n",  minExchange(values, exchange))
}
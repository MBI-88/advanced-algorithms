package main

import (
	"fmt"
	"math"
)

/*
m: Tabla de tamaño nxn almacena los minimos de multiplicaciónes escalares necesario para el
   subproducto MiMi+1...Mj. Es resultado estara en la posición m[1][n].

d: Vector que almacena las dimensiones de las matrices, siguiendo este orden: Mi = di-1 x di.

s: Diagonal de la matriz m, m se construye diagonal a diagonal. s = j-i (mij).
   s = 0, corresponde a no productos.
   s = 1, corresponde a productos de la forma MiMi+1 con un solo parentisis posible, de modo que
   mi,i+1 = di-1 x di x di+1

*/

func productOrderOptimus(d []int) int {
	n := len(d)-1
	m := make([][]int, n+1)
	step := make([][]int, n+1)
	var show func([][]int, int, int)

	for i := 0; i < n+1; i++ {
		m[i] = make([]int, n+1)
		step[i] = make([]int, n+1)
	}

	for s := 1; s < n; s++ {
		for i := 1; i < n-s+1; i++ {
			minimun := math.MaxInt 
			bestK := 0
			for k := i; k < i+s; k++ {
				cost := m[i][k] + m[k+1][i+s] + d[i-1]*d[k]*d[i+s]
				if cost < minimun {
					minimun = cost
					bestK = k
				}
			}
			m[i][i+s] = minimun
			step[i][i+s] = bestK
		}
	}

	show = func(s [][]int, i, j int) {
		if i == j {
			fmt.Printf("M%d", i)
		}else {
			k := s[i][j]
			fmt.Printf("(")
			show(s,i,k)
			show(s,k+1,j)
			fmt.Printf(")")
		}
	}
	
	show(step,1,n)
	fmt.Println()
	return m[1][n]
}

func main() {
	d := []int{13, 5, 89, 3, 34} 
	
	fmt.Printf("Costo total %d\n", productOrderOptimus(d))

}
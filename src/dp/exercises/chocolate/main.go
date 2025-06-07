package main

import (
	"fmt"
	"sort"
)

/*
Datos:
matriz MxN
Cada corte tiene un costo, depende de la recta horizontal o vertical que se corte.
Los costos denotados por cada recta:
x1,x2,x3,..., xm-1 (vertital)
y1,y2,y3,..., yn-1 (horizontal)

Ecuación:
yn-1+ X(xm-1) -> para cortes verticales, se multiplica por el total de horizontales que pasa
xm-1+ X(yn-1) -> para cortes horizontales, se multiiplica por el total de verticales que pasa


Entrada:
M-1 lineas: los valores de xm-1
N-1 lineas: los valores yn-1

Salida:
Un entero; coste total.

Requisitos:
2 <= m,n <= 1000
*/

func chocolateCut(M, N int, cuts []int) int {
	if M < 2 || M > 1000 || N < 2 || N > 1000 {
		return 0
	}
	var (
		totalCost int
		V         = cuts[:M-1]
		H         = cuts[N+1:]
		hSegment  = 1
		vSegment  = 1
	)
	sort.Sort(sort.Reverse(sort.IntSlice(V)))
	sort.Sort(sort.Reverse(sort.IntSlice(H)))

	i := 0
	j := 0
	for i < M-1 && j < N-1 {
		if V[i] > H[j] {
			totalCost += V[i] * hSegment
			i++
			vSegment++
		} else {
			totalCost += H[j] * vSegment
			j++
			hSegment++
		}
	}

	for i < M-1 {
		totalCost += V[i] * hSegment
		i++
	}
	for j < N-1 {
		totalCost += H[j] * vSegment
		j++
	}

	return totalCost
}

func main() {
	m := 6
	n := 4
	cuts := []int{2, 1, 3, 1, 4, 4, 1, 2}

	fmt.Printf("Coste total %d \n ", chocolateCut(m, n, cuts))
}

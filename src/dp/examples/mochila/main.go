package main

import "fmt"

/*
 Notas:
	La matriz que almacena los valores debe ser mas grande en sus dimensiones
	que los vectores de entrada.

	Técnica: la contribución a valores siguientes es por medio de la diagonal.
	Se observa un posicion en Yi-1,x y otra en Yi-1,x-1, esta ultima se le suma
	el beneficio y se compara con la posicion Yi-1,x.
	Si existe contribución se agrega el beneficio mas la posicion Yi-1,x-1 a la posición
	actual, Yi,x.

	Resultado:
	El resultado final estara en la posion final de la matriz, que seria la longitud de uno de
	los vectores junto a la capacidad de la mochila. R[n][M]

*/

func mochila(b, p []int, M int) int {
	n := len(b)
	R := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		R[i] = make([]int, M+1)
	}

	for i := 1; i <= n; i++ {
		for m := 1; m <= M; m++ {
			if p[i-1] <= m && b[i-1]+R[i-1][m-p[i-1]] > R[i-1][m] {
				R[i][m] = b[i-1] + R[i-1][m-p[i-1]]
			} else {
				R[i][m] = R[i-1][m]
			}
		}
	}

	return R[len(b)][M]
}

func main() {
	values := []int{10, 5, 8, 100}
	weighRs := []int{5, 2, 1, 8}
	M := 10

	fmt.Printf("Benefice %d\n", mochila(values[:], weighRs[:], M))
}

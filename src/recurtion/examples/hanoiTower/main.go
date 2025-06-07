package main

/*
Pensando primero en el caso mas pequeño, si n=1, tendríamos
un solo disco y solo habría que moverlo de la torre A a la C.
Ahora, suponiendo que para algún n ya sabemos cómo mover n−1 discos
de una torre a cualquier otra qué deberíamos hacer?
Luego de hacerse esta pregunta directamente llegamos a la conclusión de
que primero hay que mover los primeros n−1 discos a la torre B, luego el
disco n a la torre C, y posteriormente mover los n−1 discos de la torre B a
la torre C.
Podemos estar seguros que lo anterior funciona ya que los primeros n−1
discos de la torre siempre serán mas pequeños que el disco n, por lo cual
podrían colocarse libremente sobre el disco n si así lo requieran.
Se puede probar por inducción el procedimiento anterior funciona si se
hace recursivamente.
Así que nuestro algoritmo de Divide y Vencerás queda de la siguiente
manera:
Sea x la torre original, y la torre a la cual se quieren mover los discos, y
z la otra torre.
-> Para n > 1, hay que mover n−1 discos de la torre x a la z, luego mover
un disco de la torre x a la y y finalmente mover n−1 discos de la torre
z a la y.
-> Para n = 1, hay que mover el disco de la torre x a la y ignorando la
torre z.
*/

import (
	"flag"
	"fmt"
)

func hanoiTower(d int, a, b, c string) {
	if d == 1 {
		fmt.Printf("Move from %s to %s\n", a, c)
	} else {
		hanoiTower(d-1, a, c, b)
		fmt.Printf("Move from %s to %s\n", a, c)
		hanoiTower(d-1, c, b, a)
	}
}

var (
	disc *int
)

func init() {
	disc = flag.Int("discs", 4, "set total disc")
}

func main() {
	flag.Parse()
	hanoiTower(*disc, "A", "B", "C")
}

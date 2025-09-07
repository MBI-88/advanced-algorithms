package main

import (
	"fmt"
	"slices"
	"time"
)

/*
	Un nuevo experimento espacial involucra a N objetos etiquetados desde
	1 hasta N
	Se sabe de antemano que N es impar.
	Cada objeto tiene una fuerza distinta(aunque desconocida) expresada por un
	número natural.
	El objeto con la fuerza mediana es el objeto X tal que hay exactamente
	tantos objetos con una fuerza mas pequeña que X como objetos con una
	fuerza mayor que X .
	Desafortunadamente, la unica manera de comparar las fuerzas es por un
	dispositivo que, dados 3 objetos distintos determina el objeto con fuerza
	mediana tomando en cuenta solamente esos 3 objetos.

	Problema
	Escribe un programa que determine el objeto con la fuerza mediana.
	Libreria
	Dispones de una biblioteca llamada device con estas tres operaciones:
	GetN, deberá ser llamado una vez al principio sin parámetros; regresa
	el valor de N
	Med3, deberá ser llamado con tres etiquetas de objetos como parámetros; regresa la etiqueta del objeto con la fuerza media.
	Answer, deberá ser llamado una sola vez al final, con una etiqueta de
	objeto como argumento; reporta la etiqueta del objeto X
	Puedes experimentar con una biblioteca de diseño propio

	Ejemplo
	Secuencia
	2 5 4 3 1
	1 2 3 4 5
	Interacción

	1. GetN() regresa 5.
	2. Med3(1, 2, 3) regresa 3.
	3. Med3(3, 4, 1) regresa 4.
	4. Med3(4, 2, 5) regresa 4.
	5. Answer(4)
*/

type device struct {
	labelPower  map[int]int // label -> power
	powers      []int
	localLabels []int
}

func newDevice(array []int) *device {
	temp := make(map[int]int, len(array))
	for index, t := range array {
		temp[index+1] = t
	}
	return &device{
		powers:     array,
		labelPower: temp,
		localLabels: make([]int, 0,len(array) * 2),
	}
}

func (d *device) quickSelection(labels []int) int {
	var (
		space     = make([]int, 0, len(labels))
		backtrack func(int)
		mediana int
	)

	backtrack = func(i int) {
		if len(space) == 3 {
			d.localLabels = append(d.localLabels, d.media3(space[0], space[1], space[2]))
			return 
		}

		for t := i; t < len(labels); t++ {
			space = append(space, labels[t])
			backtrack(t + 1)
			space = space[:len(space)-1]
		}
	}

	backtrack(0)
	slices.Sort(d.localLabels)
	l := len(d.localLabels) 
	if l%2 == 0 {
		mediana = d.localLabels[l/2]
	}else {
		mediana = d.localLabels[(l+1)/2]
	}
	return mediana
}

func (d *device) media3(a, b, c int) int {
	fa, fb, fc := d.labelPower[a], d.labelPower[b], d.labelPower[c]
	if (fa <= fb && fb <= fc) || (fc <= fb && fb <= fa) {
		return b
	} else if (fb <= fa && fa <= fc) || (fc <= fa && fa <= fb) {
		return a
	} else {
		return c
	}
}

func (d *device) getN() int {
	return len(d.powers)
}

func (d *device) answer(label int) {
	fmt.Printf("Media global for label %d -> %d\n", label, d.labelPower[label])
}

func main() {
	array := []int{2, 5, 4, 3, 1}
	labels := make([]int, 0, len(array))
	d := newDevice(array)

	for i := range d.getN() {
		labels = append(labels, i+1)
	}
	timer := time.Now()
	d.answer(d.quickSelection(labels))
	fmt.Printf("Delay %fs total elements %d\n", time.Since(timer).Seconds(), d.getN())
}

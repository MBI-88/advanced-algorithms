package main

import (
	"flag"
	"fmt"
	"strconv"
)

func fibonacci(f int) []int {
	var (
		fibo   = make([]int, 0, f)
		fibF   func(f int)
		memory = make([]int, 0, f)
		f1     = 1
		f2     = 0
	)

	fibF = func(f int) {
		if f < 1 {
			return
		}

		for _, i := range memory {
			if f == i {
				return
			}
		}
		memory = append(memory, f)
		fn := f1 + f2
		f2 = f1
		f1 = fn
		fibo = append(fibo, f1)
		fibF(f - 1)
		return
	}
	fibF(f)
	return fibo
}

func main() {
	flag.Parse()
	input := flag.Args()[0]
	number, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Fibonacci %d\n", fibonacci(number))

}

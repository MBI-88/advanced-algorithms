package main

import (
	"flag"
	"fmt"
	"strconv"
)


func fibonacci(f int) []int {
	var (
		fibo = make([]int, 0, f) 
		f1 = 1
		f2 = 0
	)
	
    fibo = append(fibo, f1)
	for range f {
		fn := f1 + f2 // Fn = Fn-1 + Fn-2 
		f2 = f1 
		f1 = fn 
		fibo = append(fibo, f1)
	}
	return fibo
}

func main() {
	flag.Parse() 
	input := flag.Args() [0]

	fib , err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Fibonacci %d\n", fibonacci(fib))

}
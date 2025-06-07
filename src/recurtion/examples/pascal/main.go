package main

import "fmt"

func pascal(i, j int, memory [][]int) int {
	if j == 1 || i == 1 {
		return 1
	} else if memory[i][j] != 0 {
		return memory[i][j]
	}
	memory[i][j] = pascal(i-1, j-1, memory) + pascal(i-1, j, memory)
	return memory[i][j]
}

func main() {
	memory := make([][]int, 10)
	for i := range 10 {
		memory[i] = make([]int, 10)
	}
	fmt.Println(pascal(len(memory)-1, len(memory)-1, memory))
	fmt.Printf("%v\n", memory)
}

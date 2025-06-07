package main

import "fmt"

var alphabet = make([]string, 0, 21)

func combinatory(n, c int) [][]string {
	var (
		result = make([][]string, 0, c)
		temp = make([]string, 0, c)
		cad func (int)
	)

	cad = func(i1 int) {
		if len(temp) == n {
			tempCopy := make([]string, len(temp))
			copy(tempCopy, temp)
			result = append(result, tempCopy)
			return
		}

		for i := i1; i < len(alphabet[:c]); i++ {	
			temp = append(temp, alphabet[i])
			cad(i + 1)
			temp = temp[:len(temp) -1]
		}

	}

	cad(0)
	return result
}

func main() {
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	} 
	fmt.Println(combinatory(4, 5))
}
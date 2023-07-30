package main

import (
	"fmt"
)

func main() {
	fmt.Println(soma(1, 2, 3, 10, 20, 30))
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := soma(50, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(soma(10, 20))
}

func soma(a, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("O valor passou de 50")
	}
	return a + b, nil
}

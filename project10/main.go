package main

import (
	"fmt"
)

func main() {
	a := 10
	var ponteiro *int = &a
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ponteiro)
}

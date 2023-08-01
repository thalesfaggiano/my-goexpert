package main

import (
	"os"
)

func main() {
	for _, cep := range os.Args[1:] {
		println(cep)
	}
}

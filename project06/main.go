package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	faggiano := Cliente{
		Nome:  "Thales",
		Idade: 36,
		Ativo: true,
	}

	faggiano.Ativo = false

	fmt.Println(faggiano.Ativo)
}

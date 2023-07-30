package main

import "fmt"

type Endereco struct {
	Logradouro  string
	Numero      int
	Complemento string
	Cidade      string
	Estado      string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func main() {
	faggiano := Cliente{
		Nome:  "Thales",
		Idade: 36,
		Ativo: true,
	}

	fmt.Printf("Nome = %s, Idade = %d, Ativo = %t\n", faggiano.Nome, faggiano.Idade, faggiano.Ativo)
}

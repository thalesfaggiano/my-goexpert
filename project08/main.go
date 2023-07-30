package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	faggiano := Cliente{
		Nome:  "Thales",
		Idade: 36,
		Ativo: true,
	}

	faggiano.Ativo = false
	faggiano.Cidade = "São Paulo"
	faggiano.Estado = "SP"
	faggiano.Numero = 1500
	faggiano.Logradouro = "Avenida Brasil"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", faggiano.Nome, faggiano.Idade, faggiano.Ativo)
	fmt.Printf("%s, %d\n%s - %s\n", faggiano.Logradouro, faggiano.Numero, faggiano.Cidade, faggiano.Estado)
}

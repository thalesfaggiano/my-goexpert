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

func (c Cliente) desativar() {
	c.Ativo = false
	fmt.Printf("O status do cliente %s é %t", c.Nome, c.Ativo)
}

func main() {

	faggiano := Cliente{
		Nome:  "Thales",
		Idade: 36,
		Ativo: true,
	}

	faggiano.Logradouro = "Avenida Brasil"
	faggiano.Numero = 1500
	faggiano.Cidade = "São Paulo"
	faggiano.Estado = "SP"

	faggiano.desativar()

	fmt.Printf("\nO cliente %s tem %d anos e o status para atividade está em %t\n", faggiano.Nome, faggiano.Idade, faggiano.Ativo)
	fmt.Println("Endereço do cliente:")
	fmt.Printf("%s, %d\n%s - %s", faggiano.Logradouro, faggiano.Numero, faggiano.Cidade, faggiano.Estado)
}

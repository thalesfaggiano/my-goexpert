package main

import "fmt"

const a = "Hello, Faggiano"

type ID int

var (
	b bool
	c int
	d string
	e float64
	f ID
)

func main() {
	var x string = "x"
	println(a)
	fmt.Printf("O tipo da variável 'a' é: %T", a)
	println(b)
	fmt.Printf("O tipo da variável 'b' é: %T", b)
	println(c)
	fmt.Printf("O tipo da variável 'c' é: %T", c)
	println(d)
	fmt.Printf("O tipo da variável 'd' é: %T", d)
	println(e)
	fmt.Printf("O tipo da variável 'e' é: %T", e)
	println(f)
	fmt.Printf("O tipo da variável 'f' é: %T", f)
	println(x)
	fmt.Printf("O tipo da variável 'x' é: %T", x)
}

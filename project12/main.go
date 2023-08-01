package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(cep)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		defer req.Body.Close()

		println(cep)
		println(req)
	}
}

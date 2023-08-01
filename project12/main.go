package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
		}
		defer req.Body.Close()

		println("CEP = " + cep)
		println(req)

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(res))

		// bash troubleshooting:
		// diff <(go run project12/main.go 04160001) <(go run project12/main.go 04160000)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	file, err := os.Create("cep.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "File creation error: %s", err)
	}
	defer file.Close()
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Request Error: %s", err)
		}
		defer req.Body.Close()

		println("CEP = " + cep)
		println(req)

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Response Error: %s", err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", string(res))
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unmarshal Error: %s", err)
		}
		fmt.Println(data)

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento %s, Bairro %s, Localidade: %s, UF: %s\n", data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf))

		fmt.Printf("Localidade: %s\n", data.Localidade)
	}
	fmt.Println("File was created successfuly!")
}

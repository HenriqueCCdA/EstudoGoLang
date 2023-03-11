package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo cliente pelo metodo: ", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func main() {

	cliente := Cliente{
		Nome:  "Wesley",
		Email: "w@w.com",
		CPF:   12345678900,
	}

	fmt.Println(cliente)

	cliente2 := Cliente{"Mari", "m@m.com", 98765432111}

	fmt.Println("Nome: %s. Email: %s. CPF: %d", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "d@d.com",
			CPF:   98765432111,
		},
		Pais: "EUA",
	}

	fmt.Println("Nome: %s. Email: %s. Pais: %s", cliente3.Nome, cliente3.Email, cliente3.Pais)

	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	clienteJson, err := json.Marshal(cliente3)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))

	cliente4 := ClienteInternacional{}

	jsonCliente4 := `{"Nome":"Davi","Email":"d@d.com","CPF":98765432111,"pais":"EUA"}`

	json.Unmarshal([]byte(jsonCliente4), &cliente4)

	fmt.Println(cliente4)
}

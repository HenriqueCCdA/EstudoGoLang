package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

type PessoaFisica struct {
	Pessoa
	Nome      string
	Sobrenome string
	cpf       string
}

type PessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

func main() {

	p := PessoaFisica{
		Pessoa{Nome: "Tiago", Idade: 31, Status: true},
		"OU",
		"Temporim",
		"000.000.000-00",
	}

	fmt.Println(p)

}

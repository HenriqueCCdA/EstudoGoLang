package main

import "fmt"

type Document interface {
	Doc() string
}

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
	Sobrenome string
	cpf       string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func (pf PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cnpj é: %s", pf.cnpj)
}

func show(d Document) {

	switch d.(type) {
	case PessoaFisica:
		fmt.Println(d.(PessoaFisica).Sobrenome)
	case PessoaJuridica:
		fmt.Println(d.(PessoaJuridica).RazaoSocial)
	default:
		fmt.Println("tipo desconhecido")
	}

	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {

	pf := PessoaFisica{
		Pessoa{Nome: "Tiago", Idade: 31, Status: true},
		"Temporim",
		"000.000.000-00",
	}

	show(pf)

	pj := PessoaJuridica{
		Pessoa{Nome: "Aprenda Golang", Idade: 0, Status: true},
		"Temporin LTDA",
		"00.000.000/0000-00",
	}

	show(pj)

}

package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

func (c *Categoria) SetPai(pai *Categoria) {
	c.Pai = pai
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

func main() {

	cat := Categoria{Nome: "Categoria 1"}

	cat.SetPai(&Categoria{Nome: "Categoria 1"})
	if !cat.HasParent() {
		fmt.Println("Não tem pai")
	}

	p := Pessoa{Nome: "Tiago", Sobrenome: "Temporin", Idade: 31}

	fmt.Println(p)

}

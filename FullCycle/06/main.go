package main

import "fmt"

// func main() {
// 	a := 10

// 	fmt.Println(&a)

// 	var ponteiro *int = &a

// 	fmt.Println(ponteiro)
// 	fmt.Println(*ponteiro)

// 	*ponteiro = 50
// 	fmt.Println(*ponteiro)

// 	b := &a

// 	*b = 60
// 	fmt.Println(*ponteiro)

// }

type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func main() {

	variavel := 10
	abc(&variavel)

	fmt.Println(variavel)

	carro := Carro{
		Name: "Ka",
	}

	carro.andou()

	fmt.Println(carro.Name)

}

func abc(a *int) {
	*a = 200
}

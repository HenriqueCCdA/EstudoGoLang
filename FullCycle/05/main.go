package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	resultado, str := soma(10, 20)

	fmt.Println(resultado, str)

	resultado = somaTudo(10, 20, 30, 40)

	fmt.Println(resultado)

	resultado1 := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}
	fmt.Println(resultado1(10, 20)())

	carro := Carro{"Fusca"}
	carro.andar()

}

// func soma(a int, b int) (int, string) {
// 	return a + b, "somou!"
// }

func soma(a int, b int) (result int, op string) {
	result = a + b
	op = "somou!"
	return
}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}
	return resultado
}

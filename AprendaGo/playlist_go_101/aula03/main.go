package main

import "fmt"

var (
	nome string
	n1   int
	n2   int32
)

func main() {

	messagem := "Aula 03 - Go 101"
	fmt.Println(messagem)

	var b, f, s = true, 2.3, "Olá"

	fmt.Println(b, f, s)

	var total int
	total++
	fmt.Println("total:", total)

	nome = "Tiago Temporin"
	fmt.Println("Hello", nome, "!")

	var x, y = 10, 20

	fmt.Println(x, y)

	x, y = y, x

	fmt.Println(x, y)

}

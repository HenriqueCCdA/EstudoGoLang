package main

import "fmt"

func main() {

	nomes := []string{"Tiago", "Daniel", "Maria", "Marta"}

	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

	for k, nome := range nomes {
		fmt.Println(k, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	var i int

	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}

}

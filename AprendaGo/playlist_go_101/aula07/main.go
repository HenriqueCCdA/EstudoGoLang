package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

func main() {
	nomes := []string{"Tiago", "Dani", "Marcos", "Maria"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Rafael")
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Rubia")
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes2 := make([]string, 10)
	fmt.Println(nomes2, len(nomes2), cap(nomes2))

	idades := make(map[string]uint8)

	idades["Tiago"] = 31
	idades["Dani"] = 36
	idades["Maria"] = 23

	val, ok := idades["Tiago"]
	fmt.Println(val, ok)

	val, ok = idades["Lucas"]
	fmt.Println(val, ok)

	p := Pessoa{
		Nome:      "Tiago",
		Sobrenome: "Temporin",
		Idade:     31,
		Status:    true,
	}

	fmt.Println(p)

}

package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++
	case "coroa":
		coroa++
	default:
		fmt.Println("caiu em pé")
	}
}

func main() {
	a, b := 10, 13

	if a > b {
		fmt.Println("A é maior do que b")
	} else if a < b {
		fmt.Println("A é menor do que b")
	} else {
		fmt.Println("A e b são iguais")
	}

	switch {
	case a > b:
		fmt.Println("A é maior do que b")
	case a < b:
		fmt.Println("A é menor do que b")
	default:
		fmt.Println("A e b são iguais")
	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)

	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

}

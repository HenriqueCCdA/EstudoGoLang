package main

import (
	"fmt"
	"time"
)

func numeros() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func letras() {
	for i := 'a'; i < 'j'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	go numeros()
	go letras()
	time.Sleep(time.Second * 5)
	fmt.Println("Fim da execução")
}

package main

import (
	"fmt"
	"time"
)

func numeros(done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
	done <- true
}

func letras(done chan<- bool) {
	for i := 'a'; i < 'j'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Millisecond * 230)
	}
	done <- true
}

func main() {

	cn := make(chan bool)
	cl := make(chan bool)

	go numeros(cn)
	go letras(cl)

	<-cn
	<-cl

	fmt.Println("Fim da execuação.")

}

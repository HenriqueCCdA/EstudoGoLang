package main

import (
	"fmt"
	"time"
)

func main() {
	hello := make(chan string)

	go func() {
		hello <- "hello world"
	}()

	time.Sleep(time.Second * 5)
	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("default")
	}
}

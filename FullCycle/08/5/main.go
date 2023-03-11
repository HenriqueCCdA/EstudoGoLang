package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			queue <- i
			i++
		}
	}()

	for x := range queue {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}

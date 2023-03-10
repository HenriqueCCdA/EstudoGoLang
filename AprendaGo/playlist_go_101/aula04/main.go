package main

import (
	"fmt"
	"strconv"
)

func hello(nome string) {
	fmt.Println("Hello", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}

	total = a + i

	return
}

func main() {
	hello("Tiago Temporin")

	fmt.Println("total:", sum(10, 5))

	total, err := convertAndSum(10, "5")

	fmt.Println("total:", total, err)
}

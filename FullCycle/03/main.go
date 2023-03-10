package main

import (
	"03/math"
	"fmt"
)

func main() {
	resultado1 := math.Soma(1, 2)
	fmt.Printf("%v\n", resultado1)

	resultado2 := math.SomaX(1)
	fmt.Printf("%v\n", resultado2)

	fmt.Printf("%v\n", math.A)

}

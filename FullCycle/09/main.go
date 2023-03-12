package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SomaGenerica[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func Soma[T comparable](number1 T, number2 T) T {

	if number1 == number2 {
		return number1
	}

	return number2
}

func Max[T constraints.Ordered](number1 T, number2 T) T {
	if number1 > number2 {
		return number1
	}
	return number2
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func concat[T fmt.Stringer](vals []T) string {
	result := ""

	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {

	fmt.Println(concat([]MyString{"a", "b", "c"}))

	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3

	fmt.Println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))
	fmt.Println(SomaGenerica(map[string]float64{"a": 1.0, "b": 2.3, "c": 3.0}))
}

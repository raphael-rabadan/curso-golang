package main

import (
	"fmt"
	"reflect"
)

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))

	fmt.Println(reflect.TypeOf(sub))
}

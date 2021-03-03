package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(int(x) / y)
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste", 123)
	fmt.Println(fmt.Sprint(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}

	b2, _ := strconv.ParseBool("1")

	if b2 {
		fmt.Println("Verdadeiro")
	}

}

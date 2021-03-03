package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++
	y--

	fmt.Println(y, x)

	//fmt.Println(x == y++) //da erro
}

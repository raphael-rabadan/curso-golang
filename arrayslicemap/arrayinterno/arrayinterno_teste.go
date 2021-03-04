package main

import "fmt"

func teste() {

	novoArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(novoArray)

	novoSlice := novoArray[0:5]

	fmt.Println(novoArray, novoSlice)

	novoSlice[2] = 333
	fmt.Println(novoArray, novoSlice)

}

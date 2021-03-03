package main

import "fmt"

func main() {

	//var b int8 = 3
	var b byte = 3 // mesma coisa de int8

	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

}

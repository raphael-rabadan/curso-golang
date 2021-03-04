package main

import "fmt"

//func trocar(p1, p2 int) (segundo, primeiro int) { //tb funciona
func trocar(p1, p2 int) (segundo int, primeiro int) {

	segundo = p2
	primeiro = p1
	return // retorno limpo
	//return segundo, primeiro // tb funciona
}

func main() {
	x, y := trocar(2, 3)

	fmt.Println(x, y)
	fmt.Println(trocar(x, y))
	fmt.Println(trocar(trocar(x, y)))

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)

}

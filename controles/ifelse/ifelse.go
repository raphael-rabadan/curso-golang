package main

import "fmt"

func imprimirResultado(nota float64) {
	fmt.Printf("Sua nota foi %.2f e você está ", nota)
	if nota >= 7 {
		fmt.Println("Aprovado! :D")
	} else {
		fmt.Println("Reprovado! :'(")
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}

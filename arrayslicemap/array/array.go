package main

import "fmt"

func main() {
	// Sobre arrays:
	// sempre do mesmo tipo
	// tamanho fixo

	var notas [3]float64

	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	total := 0.0
	fmt.Println(notas)
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média de %.2f\n", media)
}

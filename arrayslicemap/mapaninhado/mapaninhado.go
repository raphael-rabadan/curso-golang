package main

import "fmt"

func main() {

	funcsPorLetras := map[string]map[string]float64{
		"L": {
			"Lara":  1112233.11,
			"Laila": 12121.12,
		},

		"R": {
			"Raphael": 121131.1212,
		},
	}

	delete(funcsPorLetras, "Z")

	for letra, funcs := range funcsPorLetras {
		fmt.Printf("[%s]\n", letra)

		for nome, salario := range funcs {
			fmt.Printf("%s tem salário de %.2f\n", nome, salario)
		}
	}

}

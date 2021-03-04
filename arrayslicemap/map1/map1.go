package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//Mapas precisam ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12312312311] = "Pedro"
	aprovados[99887766554] = "Maria"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados[12345678978])

	fmt.Println(aprovados)

}

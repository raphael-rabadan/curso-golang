package main

import "fmt"

// função sem parametro e sem retorno
func f1() {
	fmt.Printf("F1: Não retorno nada e nem recebo parâmetro. %s \n", "parametro do printf")
}

// função COM parâmetros e sem retorno
func f2(p1 string, p2 string) {
	fmt.Printf("F2: Parâmetros da função... %s %s \n", p1, p2)
}

// Função COM retorno de valor
func f3() string {
	return "F3: Estou retornando um valor"
}

// Função que recebe 2 parametros de mesmo valor e retorna String
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: String de retorno da %s que recebeu os parâmetros %s %s", "f4", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("param 1", "param 2")
	fmt.Println(f3())
	fmt.Println(f4("p1", "p2"))

	string1, string2 := f5()
	fmt.Printf("F5: o retorno do primeiro parametro foi [%s] e do segundo foi [%s]", string1, string2)
}

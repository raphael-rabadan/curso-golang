package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println(" linha.")

	x := 3.141516

	//Não funciona, dá erro.
	//fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	/*
		%d = inteiro
		%f = ponto flutuante
		%t = booleano
		%s = string
		%v = quase todos os tipos (não sei exatamente o que significa)
	*/
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %.1v %v %v", a, b, b, c, d)

	fmt.Printf("\n[%.10d] [%20f] [%.20f] [%20t] [%10s]", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}

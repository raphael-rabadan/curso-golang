package main

import "fmt"

func media(numeros ...float64) (media float64) {
	total := 0.0
	/*
		for i := 0; i < len(numeros); i++ {
			total += numeros[i]
		}
	*/

	for _, num := range numeros {
		total += num
	}

	if len(numeros) == 0 {
		media = 0
	} else {
		media = total / float64((len(numeros)))
	}
	return
}

func main() {

	fmt.Printf("A média foi %.2f\n", media(10.0, 10.0, 10.0, 2.0))
	fmt.Printf("A média foi %.2f\n", media(7.7, 8.1, 5.9, 9.9))
	fmt.Printf("A média foi %.2f\n", media())

}

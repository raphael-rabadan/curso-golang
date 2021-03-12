package main

import (
	"fmt"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	defer close(c)
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				break
			}
		}
	}
}

func main() {

	c := make(chan int, 30)
	go primos(1000, c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Finalizou.")

}

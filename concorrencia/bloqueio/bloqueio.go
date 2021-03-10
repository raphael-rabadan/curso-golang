package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	ch := make(chan int) // canal sem buffer

	go rotina(ch)

	fmt.Println("teste", <-ch)
	fmt.Println("Foi lido")
	fmt.Println(<-ch) // deadlock
	fmt.Println("Fim")
}

package main

import (
	"fmt"

	"github.com/raphael-rabadan/gohtml"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		gohtml.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		gohtml.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c, <-c, <-c, <-c)
}

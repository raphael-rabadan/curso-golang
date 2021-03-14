package main

import (
	"fmt"
	"time"

	"github.com/raphael-rabadan/gohtml"
)

//func oMaisRapido(url ...string) {
func oMaisRapido(url1, url2, url3 string) string {

	const tempo time.Duration = 350

	c1 := gohtml.Titulo(url1)
	c2 := gohtml.Titulo(url2)
	c3 := gohtml.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(tempo * time.Millisecond):
		return "Todos perderam"
	}
}

func main() {
	campeaoUs := oMaisRapido(
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
	)
	fmt.Println("US:", campeaoUs)

	campeaoBr := oMaisRapido(
		"https://www.google.com.br",
		"https://www.amazon.com.br",
		"https://www.youtube.com.br",
	)
	fmt.Println("BR:", campeaoBr)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!", t.Hour())
	case t.Hour() < 18:
		fmt.Println("Boa tarde!", t.Hour())
	default:
		fmt.Println("Boa noite!", t.Hour())
	}
}

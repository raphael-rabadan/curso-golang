package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {

	switch i.(type) {
	case int, int8, int16, int32, int64:
		return "inteiro"
	case uint8, uint16, uint32, uint64:
		return "inteiro sem sinal"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}

}

func main() {

	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
	fmt.Println(tipo(uint8(1)))

}

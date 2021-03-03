package main

import (
	"fmt"
	"reflect"
)

func main() {
	iTeste := int8(1)

	fmt.Println(reflect.TypeOf(iTeste))

	i := 1

	var p *int = nil //criando o ponteiro

	p = &i // pegando o endereço de referência da variável i

	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(&i))

	z := &i

	i++
	*p++
	fmt.Println(reflect.TypeOf(z), reflect.TypeOf(*z), z, *z)

}

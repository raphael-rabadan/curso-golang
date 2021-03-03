package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Printf("Valor de i2 é %d", i2)

	//números reais (float32 e float64)
	var x float32 = 49.99 //tb pode ser escrito assim: var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do tieral 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Raphael"
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da string é", len(s1))

	// string com múltiplas linhas
	s2 := `oi
	meu
	nome
	é 
	Raphael`

	fmt.Println("O tipo de s2 é", reflect.TypeOf(s2))
	fmt.Println(s2 + "!")
	fmt.Println("o tamanho da string é", len(s2))

	// char???(não existe)
	// var x char = 'b' (não existe isso)
	char := 'a' // representa o número de 'a' na tabela unicode
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}

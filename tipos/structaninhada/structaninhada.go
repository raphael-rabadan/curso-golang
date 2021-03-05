package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() (total float64) {
	total = 0.0
	for _, item := range p.itens {
		total += float64(item.qtde) * item.preco
	}
	return
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$%.2f", pedido.valorTotal())
}

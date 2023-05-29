package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	clienteID int
	itens     []item
}

func (p pedido) total() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}

	return total
}

func main() {
	pedido1 := pedido{
		clienteID: 1,
		itens: []item{
			{
				produtoID: 1,
				qtd:       3,
				preco:     5.80,
			},
			{
				produtoID: 7,
				qtd:       1,
				preco:     8.00,
			},
			{
				produtoID: 5,
				qtd:       1,
				preco:     4.50,
			},
		},
	}

	fmt.Println("o valor total do pedido deu:", pedido1.total())
}
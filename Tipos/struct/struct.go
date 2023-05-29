package main

import "fmt"

// É declarado igual em C
type produto struct {
	nome      string
	preco     float64
	descricao string
}

// A funcao recebe um dono, um estruta como receptor, exemplo  produto.desconto() . . . algo do tipo

func (p produto) desconto(desconto float64) float64 {
	return p.preco - (p.preco*desconto)/100
}

func main() {
	produto1 := produto{
		nome:      "Caderno",
		preco:     10.80,
		descricao: "Caderno capa dura",
	}

	fmt.Println(produto1,"Com desconto:",produto1.desconto(50))

	produto2 := produto{"Regua", 4.50, "regua de plastico de 15cm"}

	fmt.Println(produto2)
}

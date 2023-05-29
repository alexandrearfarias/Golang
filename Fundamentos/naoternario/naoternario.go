package main

import "fmt"

//FORMA DE REPRESENTAR UM TERNARIO
func obterV(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

// return nota>=6 ? "Aprovado": "Reprovado"

func main() {
	fmt.Println(obterV(7))
}

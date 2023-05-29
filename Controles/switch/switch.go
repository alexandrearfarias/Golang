package main

import (
	"fmt"
)

func obterNota(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough // palavra reservda pra pular um dos casos
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5, 4:
		return "C"
	case 3, 2, 1:
		return "E"
	default:
		return "nota invalida"
	}
}

func desafio(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "nota A"
	case n >= 8 && n < 9:
		return "nota B"
	case n >= 5 && n < 8:
		return "nota C"
	case n >= 3 && n < 5:
		return "nota D"
	default:
		return "nota E"
	}
}

func main() {

	fmt.Println(obterNota(10))
	fmt.Println(obterNota(12))
	fmt.Println(obterNota(4))
	fmt.Println(obterNota(7))

	fmt.Println(desafio(10))
	fmt.Println(desafio(2))
	fmt.Println(desafio(5.5))
}

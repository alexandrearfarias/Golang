package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//o make tem 3 paramentos: o tipo de slice, o tamanho e sua capacidade(tamanho do array interno)
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	//Adiciona um elementos ao slice
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// Quando o array interno chega na sua capacidade maxima, caso seja adicionado mais um elemento, ele dobra de tamanho
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}

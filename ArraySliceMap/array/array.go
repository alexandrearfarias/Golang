package main

import "fmt"

// array é fixo e homogeneno(aceita só um tipo)
func main() {
	var notas [3]float64

	fmt.Println(notas)

	numeros := [...]int{1, 2, 3, 4, 5, 6}

	// No for range sempre vai precisar de duas variaveis, o index e o elemento do array
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	//O underline é usado para ignorar alguma coisa obrigatória do código
	for _, num := range numeros {
		fmt.Printf("%d-", num)
	}
}

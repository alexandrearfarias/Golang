package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func imprimirAprovados(aprovados ...string) {
	for i, aprovado := range aprovados {
		fmt.Printf("%d ) %s \n", i, aprovado)
	}
}

func main() {
	fmt.Printf("Media: %2.f \n", media(7.7, 8.9, 9, 4.5))

	fmt.Println("Lista de aprovados")
	lista := []string{"joao", "maria", "pedro"}

	imprimirAprovados(lista...)
}

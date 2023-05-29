package main

import "fmt"

//Quando tem um tipo fora dos paramentos da função, necessariamente ela precisa de um return
func soma(a int, b int) int {
	return a + b
}

func imprimirV(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := soma(123, 5)

	imprimirV(resultado)
}

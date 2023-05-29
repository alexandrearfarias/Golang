package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.14
	var raio = 3.6
	
	//O ':=' cria e define a variavel na mesma lapada 
	area := PI * m.Pow(raio,2)

	//Sempre usar a variavel, depois de declarada
	fmt.Printf("a area do redondo é: %.2f\n", area)

	//Criar variaveis em uma linha só
	a, b, c := true, 2, "vish"
	fmt.Println(a,b,c)
}
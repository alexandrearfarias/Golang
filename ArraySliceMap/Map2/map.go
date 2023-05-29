package main

import "fmt"

func main() {
	lista := map[string]float64{
		"Jose":    10.50,
		"Maria":   2.80,
		"Pedro":   7.90,
		"Sampaio": 17.10,
	}

	lista["Samuel Silva"] = 5.90

	for nome, salario := range lista {
		fmt.Printf("%s; salario: %.2f\n", nome,salario)
	}

}
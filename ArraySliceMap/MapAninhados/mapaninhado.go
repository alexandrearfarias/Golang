package main

import "fmt"

func main() {
	porLetra := map[string]map[string]float64{
		"S": {
			"Samuel Silva": 24.07,
			"Samara":       11.50,
		},
		"F": {
			"Filkard":  800.60,
			"Fcardoli": 56.80,
		},
	}

	for letra, funcionarios := range porLetra {
		fmt.Println(letra, funcionarios)
	}
}

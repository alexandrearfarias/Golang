package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.99):
		return "Nota A"
	case n.entre(7.0, 8.99):
		return "Nota B"
	case n.entre(5.0, 6.99):
		return "Nota C"
	case n.entre(3.0, 4.99):
		return "Nota C"
	default:
		return "Nota invalida"
	}
}

func main() {

	fmt.Println(notaConceito(6.5))
	fmt.Println(notaConceito(7.8))
	fmt.Println(notaConceito(9.5))
	fmt.Println(notaConceito(2.1))

}
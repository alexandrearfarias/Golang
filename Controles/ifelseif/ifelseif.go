package main

import "fmt"

func obtervalor(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "Aluno nota A"
	} else if nota >= 8 && nota < 9 {
		return "Aluno nota B"
	} else if nota >= 5 && nota < 8 {
		return "Aluno nota C"
	} else if nota >= 3 && nota < 5 {
		return "Aluno nota D"
	} else {
		return "Aluno nota E"
	}
}

func main() {
	fmt.Println(obtervalor(10))
	fmt.Println(obtervalor(6.5))
	fmt.Println(obtervalor(2))
}

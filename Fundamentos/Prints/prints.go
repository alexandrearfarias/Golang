package main

import "fmt"

func main() {
	fmt.Print("Printando")
	fmt.Print("Na mesma linha")

	fmt.Println("\nPrintando em linhas diferentes")

	/*O priint ln aceita apenas o formato string
	A função Sprint transforma no tipo string
	*/
	x := 3.14
	xs := fmt.Sprint(x)
	fmt.Println("print ln so aceita strings, veja:" + xs)

	a, b, c, d := 1, 1.9, false, "opa"

	//Os ganchos para cada tipo de variavel
	fmt.Printf("\n %d %.1f %t %s", a, b, c, d)
}

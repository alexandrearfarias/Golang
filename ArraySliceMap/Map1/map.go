package main

import "fmt"

func main() {
	//var aprovados map[int]string
	aprovados := make(map[int]string) // a chave e o valor do map podem ser diferentes

	aprovados[123] = "Miguel"
	aprovados[456] = "Jefferson"
	aprovados[789] = "caminhoes"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados{
		fmt.Printf("%s CPF(%d) \n", nome,cpf)
	}

	// pegar o valor de uma chave do map
	fmt.Println(aprovados[123])

	// deletar uma chave do map
	delete(aprovados,789)
	fmt.Println(aprovados[789])
}
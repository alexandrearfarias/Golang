package main

import "fmt"

func main() {
	a := 1

	var p *int = nil //pode criar ponteiros nulos(vazios),diferentemente das variaveis normais

	p = &a // o &a é para atribuir o endereço de memoria da variavel ao ponteiro P

	*p++ // pra ter acesso ao dado endereçado no ponteiro usa *p

	fmt.Println(p,"<- endereço", *p, a, "<-valor", &a, "<-endereço")
}